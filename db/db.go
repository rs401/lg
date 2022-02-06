// Package db provides utilities to configure and connect to a database
package db

import (
	"log"

	"github.com/rs401/lg/auth/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectionError custom error type
type ConnectionError struct{}

// Error implements the error interface
func (ce *ConnectionError) Error() string {
	return "error connecting to database"
}

// Connection interface defines a method for retrieving the *gorm.DB
type Connection interface {
	DB() *gorm.DB
}

type conn struct {
	db *gorm.DB
}

// DB implements the Connection interface
func (c *conn) DB() *gorm.DB {
	return c.db
}

// NewConnection takes a db.Config and returns a Connection and an error
func NewConnection(cfg Config) (Connection, error) {
	dbc, err := gorm.Open(postgres.Open(cfg.ConnStr()), &gorm.Config{})
	if err != nil {
		log.Printf("Error, could not connect to database: %v", err)
		return nil, &ConnectionError{}
	}
	// I guess do this here
	dbc.AutoMigrate(&models.User{})
	return &conn{db: dbc}, nil
}
