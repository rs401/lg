package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/rs401/lg/auth/repository"
	"github.com/rs401/lg/auth/service"
	"github.com/rs401/lg/db"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file (production?): %v\n", err)
	}
}

func main() {
	// Get our config and a new connection
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Get a usersRepository
	usersRepository := repository.NewUsersRepository(conn)
	// Get an authService
	authService := service.NewAuthService(usersRepository)
	// Register authService with RPC
	err = rpc.Register(authService)
	if err != nil {
		log.Fatalf("Error registering auth service: %v\n", err)
	}
	// RPC Handle HTTP
	rpc.HandleHTTP()
	// Get auth service port from environment
	port, err := strconv.Atoi(os.Getenv("AUTHSVC_PORT"))
	if err != nil {
		log.Fatalf("Error getting auth service port: %v\n", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	log.Printf("Auth service running on port: :%d\n", port)
	log.Fatal(http.Serve(lis, nil))
}
