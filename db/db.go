package db

import (
	"log"

	"github.com/rs401/lg/auth/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connection interface {
	DB() *gorm.DB
}

type conn struct {
	db *gorm.DB
}

func (c *conn) DB() *gorm.DB {
	return c.db
}

func NewConnection(cfg Config) (Connection, error) {
	dbc, err := gorm.Open(postgres.Open(cfg.ConnStr()), &gorm.Config{})
	if err != nil {
		log.Printf("Error, could not connect to database: %v", err)
		return nil, err
	}
	// I guess do this here
	dbc.AutoMigrate(&models.User{})
	return &conn{db: dbc}, nil
}
