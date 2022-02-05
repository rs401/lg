package repository

import (
	"errors"

	"github.com/rs401/lg/auth/models"
	"github.com/rs401/lg/db"
	"github.com/rs401/lg/validation"
	"gorm.io/gorm"
)

var ErrorBadID error = errors.New("bad id")

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

func NewUsersRepository(conn db.Connection) UsersRepository {
	return &usersRepository{db: conn.DB()}
}

func (r *usersRepository) Save(user *models.User) error {
	return r.db.Create(&user).Error
}

func (r *usersRepository) GetById(id uint) (*models.User, error) {
	var user models.User
	result := r.db.Where("ID = ?", id).First(&user)
	return &user, result.Error
}

func (r *usersRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	result := r.db.Where("email = ?", email).Find(&user)
	return &user, result.Error
}

func (r *usersRepository) GetAll() ([]*models.User, error) {
	var ul models.GetUsersResponse
	result := r.db.Find(&ul.Users)
	return ul.Users, result.Error
}

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

func (r *usersRepository) Delete(id uint) error {
	var user models.User
	r.db.Find(&user, id)
	if user.ID == 0 {
		return ErrorBadID
	}
	return r.db.Delete(&user).Error
}

func (r *usersRepository) DeleteAll() error {
	// return r.db.Where("1 = 1").Delete(&models.User{}).Error
	return r.db.Exec("DELETE FROM users").Error
}
