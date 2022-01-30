package service

import (
	"context"
	"strings"
	"time"

	"github.com/rs401/lg/auth/models"
	"github.com/rs401/lg/auth/repository"
	"github.com/rs401/lg/validation"
	"golang.org/x/crypto/bcrypt"
)

type AuthSvc interface {
	SignUp(*models.SignUpRequest, *models.User) error
	SignIn(*models.SignInRequest, *models.User) error
	GetUser(*models.GetUserRequest, *models.User) error
	ListUsers(context.Context, *models.UserList) error
	UpdateUser(*models.User, *models.User) error
	DeleteUser(*models.GetUserRequest, *models.GetUserRequest) error
}

type AuthService struct {
	usersRepository repository.UsersRepository
}

func NewAuthService(usersRepository repository.UsersRepository) AuthSvc {
	return &AuthService{usersRepository: usersRepository}
}

func (as *AuthService) SignUp(req *models.SignUpRequest, res *models.User) error {
	err := validation.IsValidSignUp(req)
	if err != nil {
		return err
	}
	exists, err := as.usersRepository.GetByEmail(req.Email)
	if err != nil {
		return err
	}
	if exists.Name != "" {
		return validation.ErrEmailExists
	}

	if exists.Name == "" {
		// user := new(models.User)
		res.Name = strings.TrimSpace(req.Name)
		res.Email = validation.NormalizeEmail(req.Email)
		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		res.Password = hash

		err = as.usersRepository.Save(res)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate") {
				if strings.Contains(err.Error(), "name") {
					return validation.ErrNameExists
				}
				if strings.Contains(err.Error(), "email") {
					return validation.ErrEmailExists
				}
			}
			return err
		}
		// res = &user
		return nil
	}

	return err

}

func (as *AuthService) SignIn(req *models.SignInRequest, res *models.User) error {
	user, err := as.usersRepository.GetByEmail(req.Email)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword(user.Password, []byte(req.Password))
	if err != nil {
		return err
	}
	res.ID = user.ID
	res.Name = user.Name
	res.Email = user.Email
	return nil
}

func (as *AuthService) GetUser(req *models.GetUserRequest, res *models.User) error {
	user, err := as.usersRepository.GetById(req.Id)
	if err != nil {
		return err
	}
	res.Name = user.Name
	res.Email = user.Email
	// res = user

	return nil
}

func (as *AuthService) ListUsers(ctx context.Context, res *models.UserList) error {
	users, err := as.usersRepository.GetAll()
	if err != nil {
		return err
	}

	res.Users = append(res.Users, users...)

	return nil
}

func (as *AuthService) UpdateUser(req *models.User, res *models.User) error {
	// Verify user exists
	user, err := as.usersRepository.GetById(uint(req.ID))
	if err != nil {
		return err
	}
	if user == nil {
		return validation.ErrNotFound
	}

	// Validate user name not empty
	if validation.IsEmptyString(req.Name) {
		return validation.ErrEmptyName
	}

	// Validate user email not empty
	if validation.IsEmptyString(req.Email) {
		return validation.ErrEmptyEmail
	}

	// Validate user email is email
	if !validation.IsValidEmail(req.Email) {
		return validation.ErrInvalidEmail
	}

	// Update the user record
	res.Name = req.Name
	res.Email = req.Email
	res.UpdatedAt = time.Now()

	err = as.usersRepository.Update(res)
	if err != nil {
		return err
	}
	return nil

}

func (as *AuthService) DeleteUser(req, res *models.GetUserRequest) error {
	err := as.usersRepository.Delete(uint(req.Id))
	if err != nil {
		return err
	}
	return nil
}
