package server

import (
	"context"
	"errors"
	"github.com/gorilla/mux"
	"github.com/lelandaure/rest-go-ws/database"
	"github.com/lelandaure/rest-go-ws/repository"
	"log"
	"net/http"
)

type Config struct {
	Port        string
	JWTSecret   string
	DatabaseUrl string
}

type Server interface {
	Config() *Config
}

type Broker struct {
	config *Config
	router *mux.Router
}

func (b *Broker) Config() *Config {
	return b.config
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("port is required")
	}
	if config.JWTSecret == "" {
		return nil, errors.New("JWTSecret is required")
	}
	if config.DatabaseUrl == "" {
		return nil, errors.New("DatabaseUrl is required")
	}

	return &Broker{
			config: config,
			router: mux.NewRouter(),
		},
		nil
}

func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	b.router = mux.NewRouter()
	binder(b, b.router)
	repo, err := database.NewPostgresRepository(b.config.DatabaseUrl)
	if err != nil {
		log.Fatalln(err)
	}
	repository.SetRepository(repo)
	log.Println("Starting server on port", b.config.Port)
	if err := http.ListenAndServe(b.config.Port, b.router); err != nil {
		log.Fatal("ListenAndServerError:", err)
	}
}
