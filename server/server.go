package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/carlosescorche/qrlist/database"
	"github.com/gorilla/mux"
)

type Config struct {
	Port         string
	DatabaseURL  string
	DatabaseName string
}

type Server struct {
	Config *Config
	Router *mux.Router
	Hub    *Hub
}

type bindRoutes func(s *Server)

func NewServer(ctx context.Context, config *Config) (*Server, error) {
	if config.Port == "" {
		return nil, fmt.Errorf("Port is required")
	}

	if config.DatabaseURL == "" {
		return nil, fmt.Errorf("Database URL is required")
	}

	if config.DatabaseName == "" {
		return nil, fmt.Errorf("Database name is required")
	}

	s := &Server{
		Config: config,
		Router: mux.NewRouter(),
		Hub:    NewHub(),
	}

	return s, nil
}

func (s *Server) Start(binder bindRoutes) {
	database.Connect(s.Config.DatabaseURL, s.Config.DatabaseName)

	defer func() {
		database.Close()
	}()

	// Sets routes
	binder(s)

	go s.Hub.Run()

	log.Printf("Server running on port %v", s.Config.Port)
	if err := http.ListenAndServe(s.Config.Port, s.Router); err != nil {
		log.Printf("Error starting server: %v", err)
	}
}
