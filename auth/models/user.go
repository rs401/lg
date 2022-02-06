// Package models provides data structures
package models

import (
	"time"
)

// User model for user
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"<-;unique;not null"`
	Email     string    `json:"email" gorm:"<-;unique;not null"`
	Password  []byte    `json:"-"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
