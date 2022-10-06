package handler

import (
	"encoding/json"
	"github.com/lelandaure/rest-go-ws/models"
	"github.com/lelandaure/rest-go-ws/repository"
	"github.com/lelandaure/rest-go-ws/server"
	"github.com/segmentio/ksuid"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
)

const (
	HashCost = 8
)

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

func SignUpHandler(s server.Server) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		request := SignUpRequestBuilder(w, r.Body)
		hashedPassword := hashedPasswordBuilder(w, request)
		randomId, err := ksuid.NewRandom()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var user = models.User{Id: randomId.String(), Email: request.Email, Password: string(hashedPassword)}

		err = repository.InsertUser(r.Context(), &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().
			Set("Content-Type", "application/json")

		err = json.NewEncoder(w).Encode(SignUpResponse{
			Id:    user.Id,
			Email: user.Email,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

}

func SignUpRequestBuilder(w http.ResponseWriter, body io.ReadCloser) SignUpRequest {
	var request SignUpRequest
	err := json.NewDecoder(body).
		Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return SignUpRequest{}
	}
	return request
}

func hashedPasswordBuilder(w http.ResponseWriter, request SignUpRequest) []byte {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), HashCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}
	return hashedPassword
}
