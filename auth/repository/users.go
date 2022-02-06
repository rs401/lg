// Package repository provides methods to interact with the database
package repository

import (
	"errors"

	"github.com/rs401/lg/auth/models"
	"github.com/rs401/lg/db"
	"github.com/rs401/lg/validation"
	"gorm.io/gorm"
)

// ErrorBadID custom error
var ErrorBadID error = errors.New("bad id")

// UsersRepository interface defines methods for interacting with the database
type UsersRepository interface {
	Save(user *models.User) error
	GetById(id uint) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetAll() ([]*models.User, error)
	Update(user *models.User) error
	Delete(id uint) error
}

type usersRepository struct {
	db *gorm.DB
}

// NewUsersRepository takes a db.Connection and returns a UsersRepository
func NewUsersRepository(conn db.Connection) UsersRepository {
	return &usersRepository{db: conn.DB()}
}

// Save takes a *models.User and saves to the database
func (r *usersRepository) Save(user *models.User) error {
	return r.db.Create(&user).Error
}

// GetById takes a uint id and returns a *models.User and an error
func (r *usersRepository) GetById(id uint) (*models.User, error) {
	var user models.User
	result := r.db.Where("ID = ?", id).First(&user)
	return &user, result.Error
}

// GetByEmail takes an email string and returns a *models.User and an error
func (r *usersRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	result := r.db.Where("email = ?", email).Find(&user)
	return &user, result.Error
}

// GetAll returns a slice of *models.User and an error
func (r *usersRepository) GetAll() ([]*models.User, error) {
	var ul models.GetUsersResponse
	result := r.db.Find(&ul.Users)
	return ul.Users, result.Error
}

// Update takes a *models.User and updates the record in the database
func (r *usersRepository) Update(user *models.User) error {
	var tmpUser = new(models.User)
	r.db.Find(&tmpUser, user.ID)
	if tmpUser.Name != user.Name && !validation.IsEmptyString(user.Name) {
		tmpUser.Name = user.Name
	}
	if tmpUser.Email != user.Email && validation.IsValidEmail(user.Email) {
		tmpUser.Email = user.Email
	}
	return r.db.Save(&tmpUser).Error
}

// Delete takes a uint id and deletes a record from the database
func (r *usersRepository) Delete(id uint) error {
	var user models.User
	r.db.Find(&user, id)
	if user.ID == 0 {
		return ErrorBadID
	}
	return r.db.Delete(&user).Error
}

// DeleteAll deletes all records from the database
func (r *usersRepository) DeleteAll() error {
	// return r.db.Where("1 = 1").Delete(&models.User{}).Error
	return r.db.Exec("DELETE FROM users").Error
}
