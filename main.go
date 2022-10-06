package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/lelandaure/rest-go-ws/handler"
	"github.com/lelandaure/rest-go-ws/server"
	"log"
	"net/http"
	"os"
)

var err = godotenv.Load(".env")

var PORT = os.Getenv("PORT")
var JwtSecret = os.Getenv("JWT_SECRET")
var DatabaseUrl = os.Getenv("DATABASE_URL")

func main() {
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	myServer, err := server.NewServer(
		context.Background(),
		&server.Config{Port: PORT, JWTSecret: JwtSecret, DatabaseUrl: DatabaseUrl},
	)

	if err != nil {
		log.Fatal("ListenAndServerError: ", err)
	}

	myServer.Start(BindRoutes)
}

func BindRoutes(s server.Server, r *mux.Router) {
	r.HandleFunc("/", handler.HomeHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/signup", handler.SignUpHandler(s)).Methods(http.MethodPost)
}
