package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs401/lg/api/handlers"
)

func SetupRoutes(r *mux.Router, hndlrs handlers.AuthHandlers) {
	r.HandleFunc("/api/", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(map[string]string{"heartbeat": "alive"})
	})

	r.HandleFunc("/api/signup/", hndlrs.SignUp).Methods("POST")
	r.HandleFunc("/api/signin/", hndlrs.SignIn).Methods("POST")
	r.HandleFunc("/api/user/", hndlrs.GetUsers).Methods("GET")
	r.HandleFunc("/api/user/{id:[0-9]+}/", hndlrs.GetUsers).Methods("GET")
	r.HandleFunc("/api/user/{id:[0-9]+}/", hndlrs.UpdateUser).Methods("PUT")
	r.HandleFunc("/api/user/{id:[0-9]+}/", hndlrs.DeleteUser).Methods("DELETE")
}
