package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/rs401/lg/api/authclient"
	"github.com/rs401/lg/api/tokenutils"
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

// NewAuthHandlers Creates new auth handlers
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

	// Might as well create tokens so they don't have to log in
	tokens, err := tokenutils.CreateToken(result.ID)
	if err != nil {
		log.Printf("Unable to create tokens after signup: %v\n", err)
	}
	if tokens != nil {
		// Set access token in auth header
		w.Header().Set("Authorization", tokens.AccessToken)
		// Set refresh token in cookie
		http.SetCookie(w,
			&http.Cookie{
				Name:    "refresh_token",
				Value:   tokens.RefreshToken,
				Expires: time.Now().Add(time.Hour * 24),
			})
	}
	// Let them know
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

func (ah *authHandlers) SignIn(w http.ResponseWriter, r *http.Request) {
	var signReq models.SignInRequest
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&signReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "json formating decode error"})
		return
	}
	err = ah.authSvcClient.SignIn(&signReq, &user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	if user.Name == "" {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "user not found"})
		return
	}
	// All good, create tokens
	tokens, err := tokenutils.CreateToken(user.ID)
	if err != nil {
		log.Printf("Unable to create tokens after signin: %v\n", err)
	}
	if tokens != nil {
		// Set access token in auth header
		w.Header().Set("Authorization", tokens.AccessToken)
		// Set refresh token in cookie
		http.SetCookie(w,
			&http.Cookie{
				Name:    "refresh_token",
				Value:   tokens.RefreshToken,
				Expires: time.Now().Add(time.Hour * 24),
			})
	}
	// Let them know
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (ah *authHandlers) UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func (ah *authHandlers) GetUser(w http.ResponseWriter, r *http.Request) {

}

func (ah *authHandlers) GetUsers(w http.ResponseWriter, r *http.Request) {

}

func (ah *authHandlers) DeleteUser(w http.ResponseWriter, r *http.Request) {

}
