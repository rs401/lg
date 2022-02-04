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
	AuthSvcHost string
	AuthSvcPort int
	APIPort     int
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file (production?): %v\n", err)
	}
	AuthSvcHost = os.Getenv("AUTHSVC_HOST")
	AuthSvcPort, err = strconv.Atoi(os.Getenv("AUTHSVC_PORT"))
	if err != nil {
		log.Fatalf("Error converting AUTHSVC_PORT to int")
	}
	APIPort, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		log.Fatalf("Error converting API_PORT to int")
	}
}
func main() {
	// Client dial server
	client, err := rpc.DialHTTP("tcp", fmt.Sprintf("%s:%d", AuthSvcHost, AuthSvcPort))
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
	log.Printf("Listening on port :%d\n", APIPort)
	err = http.ListenAndServe(fmt.Sprintf(":%d", APIPort), router)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
}
