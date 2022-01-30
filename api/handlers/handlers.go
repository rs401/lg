package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rs401/lg/api/authclient"
	"github.com/rs401/lg/auth/models"
)

type AuthHandlers interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	SignIn(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type authHandlers struct {
	authSvcClient authclient.AuthSvcClient
}

func NewAuthHandlers(authSvcClient authclient.AuthSvcClient) AuthHandlers {
	return &authHandlers{authSvcClient: authSvcClient}
}

func (ah *authHandlers) SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.SignUpRequest
	var result models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "json formating decode error"})
		return
	}

	err = ah.authSvcClient.SignUp(&user, &result)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

func (ah *authHandlers) SignIn(w http.ResponseWriter, r *http.Request) {

}

func (ah *authHandlers) UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func (ah *authHandlers) GetUser(w http.ResponseWriter, r *http.Request) {

}

func (ah *authHandlers) GetUsers(w http.ResponseWriter, r *http.Request) {

}

func (ah *authHandlers) DeleteUser(w http.ResponseWriter, r *http.Request) {

}
