package main

import (
	"context"
	"os"

	"github.com/carlosescorche/qrlist-api/handlers"
	"github.com/carlosescorche/qrlist-api/middlewares"
	"github.com/carlosescorche/qrlist-api/server"
)

func routes(s *server.Server) {
	r := s.Router
	r.Use(middlewares.MiddlewareCors)

	r.HandleFunc("/list", handlers.HandlerListCreate(s)).Methods("POST")
	r.HandleFunc("/list/{id}", handlers.HandlerListGet(s)).Methods("GET")
	r.HandleFunc("/list/{id}/subscriptions", handlers.HandlerListSubscriptions(s)).Methods("GET")
	r.HandleFunc("/list/{id}/signup", handlers.HandlerListSignup(s)).Methods("POST")
	r.HandleFunc("/list/{id}/observe", handlers.HandlerListObserve(s))
	r.HandleFunc("/subscription/{id}", handlers.HandlerSubscriptionUpdate(s)).Methods("POST")
	r.HandleFunc("/subscription/{id}", handlers.HandlerSubscriptionGet(s)).Methods("GET")
}

func main() {
	s, err := server.NewServer(context.Background(), &server.Config{
		Port:         os.Getenv("PORT"),
		DatabaseURL:  os.Getenv("DATABASE_URL"),
		DatabaseName: os.Getenv("DATABASE_NAME"),
	})

	if err != nil {
		panic(err)
	}

	s.Start(routes)
}
