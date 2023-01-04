package server

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type Client struct {
	listId string
	hub    Hub
	conn   *websocket.Conn
	send   chan []byte
}

func NewClient(listId string, hub *Hub, socket *websocket.Conn) *Client {
	return &Client{
		listId: listId,
		hub:    *hub,
		conn:   socket,
		send:   make(chan []byte),
	}
}

func (c *Client) Write() {
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.conn.WriteMessage(websocket.TextMessage, message)
		}
	}
}

type Hub struct {
	Clients    map[*Client]bool
	Register   chan *Client
	Unregister chan *Client
	Mutex      *sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[*Client]bool),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Mutex:      &sync.Mutex{},
	}
}

func (hub *Hub) Run() {
	for {
		select {
		case client := <-hub.Register:
			hub.onConnect(client)
		case client := <-hub.Unregister:
			hub.onDisconnect(client)
		}
	}
}

func (hub *Hub) Broadcast(message interface{}, listId string) {
	data, _ := json.Marshal(message)

	for client := range hub.Clients {
		if client.listId == listId {
			client.send <- data
		}
	}
}

func (hub *Hub) onConnect(client *Client) {
	log.Println("Client connected", client)

	hub.Mutex.Lock()
	defer hub.Mutex.Unlock()

	hub.Clients[client] = true
}

func (hub *Hub) onDisconnect(client *Client) {
	log.Println("Client disconnected", client)

	hub.Mutex.Lock()
	defer hub.Mutex.Unlock()

	if _, ok := hub.Clients[client]; ok {
		delete(hub.Clients, client)
		client.conn.Close()
		close(client.send)
	}
}
