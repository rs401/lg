// Package api provides HTTP endpoints
package main

import (
	"fmt"
	"log"
	"net/http"
	"net/rpc"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs401/lg/api/authclient"
	"github.com/rs401/lg/api/handlers"
	"github.com/rs401/lg/api/middlewares"
	"github.com/rs401/lg/api/routes"
)

var (
	authSvcHost string
	authSvcPort int
	apiPort     int
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file (production?): %v\n", err)
	}
	authSvcHost = os.Getenv("AUTHSVC_HOST")
	authSvcPort, err = strconv.Atoi(os.Getenv("AUTHSVC_PORT"))
	if err != nil {
		log.Fatalf("Error converting AUTHSVC_PORT to int")
	}
	apiPort, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		log.Fatalf("Error converting API_PORT to int")
	}
}
func main() {
	// Client dial server
	client, err := rpc.DialHTTP("tcp", fmt.Sprintf("%s:%d", authSvcHost, authSvcPort))
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// Initialize client
	// Pass this to handlers
	authSvcClient := authclient.NewAuthServiceClient(client)
	// Setup handlers
	hndlrs := handlers.NewAuthHandlers(authSvcClient)
	// Setup router
	router := mux.NewRouter()
	// Setup routes
	routes.SetupRoutes(router, hndlrs)
	// Setup middlewares
	middlewares.SetupMiddleWares(router)
	// Listen
	log.Printf("Listening on port :%d\n", apiPort)
	err = http.ListenAndServe(fmt.Sprintf(":%d", apiPort), router)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
}
