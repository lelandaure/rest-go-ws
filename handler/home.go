package handler

import (
	"encoding/json"
	"github.com/lelandaure/rest-go-ws/server"
	"net/http"
)

type HomeResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func HomeHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().
			Set("Content-Type", "application/json")

		err := json.NewEncoder(w).
			Encode(HomeResponse{"Welcome to api GO", true})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}
