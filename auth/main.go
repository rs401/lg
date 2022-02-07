// Package auth provides RPC methods to be called
package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs401/lg/auth/repository"
	"github.com/rs401/lg/auth/service"
	"github.com/rs401/lg/db"
)

var (
	dbRetries       uint
	dbMaxRetries    uint
	dbRetryDuration time.Duration
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file (production?): %v\n", err)
	}
	dbRetries = 0
	dbMaxRetries = 10
	dbRetryDuration = time.Second * 3
}

func main() {
	// Get our config and a new connection
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	for err != nil && dbRetries <= dbMaxRetries {
		_, ok := err.(*db.ConnectionError)
		if ok {
			// connection error, db might not be up
			log.Printf("Error connecting to db, retrying in %f seconds...\n", dbRetryDuration.Seconds())
			time.Sleep(dbRetryDuration)
			dbRetries++
			conn, err = db.NewConnection(cfg)
		} else {
			log.Fatalf("Error connecting to database: %v", err)
		}
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
