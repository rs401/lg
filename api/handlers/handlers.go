// Package handlers provides handlerfuncs
package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs401/lg/api/authclient"
	"github.com/rs401/lg/api/tokenutils"
	"github.com/rs401/lg/auth/models"
)

// AuthHandlers interface defining HandlerFuncs
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

// NewAuthHandlers takes an authclient.AuthSvcClient and returns an AuthHandlers
func NewAuthHandlers(authSvcClient authclient.AuthSvcClient) AuthHandlers {
	return &authHandlers{authSvcClient: authSvcClient}
}

// SignUp handles calling the Client.SignUp method
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

// SignIn handles calling the Client.SignIn method
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
		w.Header().Set("Authorization", fmt.Sprintf("Bearer %v", tokens.AccessToken))
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

// UpdateUser handles calling the Client.UpdateUser method
func (ah *authHandlers) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// TODO: Verify authenticated user is the user being updated
	vars := mux.Vars(r)
	var user = new(models.User)
	var result = new(models.User)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "bad path"})
		return
	}
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "json formating decode error"})
		return
	}
	if user.ID != uint(id) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "bad request, id error"})
		return
	}
	err = ah.authSvcClient.UpdateUser(user, result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

// GetUser handles calling the Client.GetUser method
func (ah *authHandlers) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var req models.GetUserRequest
	var user models.User
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "bad path"})
		return
	}
	req.Id = uint(id)
	err = ah.authSvcClient.GetUser(&req, &user)
	if err != nil {
		log.Printf("Error calling GetUser: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(user)
}

// GetUsers handles calling the Client.ListUsers method
func (ah *authHandlers) GetUsers(w http.ResponseWriter, r *http.Request) {
	var users models.GetUsersResponse
	err := ah.authSvcClient.ListUsers("trash", &users)
	if err != nil {
		log.Printf("Error calling ListUsers: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users.Users)
}

// DeleteUser handles calling the Client.DeleteUser method
func (ah *authHandlers) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// TODO: Verify authenticated user is the user being deleted
	vars := mux.Vars(r)
	var req, res models.GetUserRequest
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "bad path"})
		return
	}
	req.Id = uint(id)
	err = ah.authSvcClient.DeleteUser(&req, &res)
	if err != nil {
		log.Printf("Error calling DeleteUser: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"delete": "success"})
}
