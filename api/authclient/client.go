package authclient

import (
	"context"

	"github.com/rs401/lg/auth/models"
)

var (
	AuthSvcHost string
	AuthSvcPort int
)

type AuthClientI interface {
	SignUp(*models.User, *models.User) error
	GetUser(*models.GetUserRequest, *models.User) error
	ListUsers(context.Context, *models.UserList) error
	UpdateUser(*models.User, *models.User) error
	DeleteUser(*models.GetUserRequest, *models.GetUserRequest) error
}

type AuthClient struct {
}
