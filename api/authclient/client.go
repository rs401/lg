package authclient

import (
	"context"
	"net/rpc"

	"github.com/rs401/lg/auth/models"
)

type AuthSvcClient interface {
	SignUp(req *models.SignUpRequest, res *models.User) error
	SignIn(req *models.SignInRequest, res *models.User) error
	GetUser(req *models.GetUserRequest, res *models.User) error
	ListUsers(req context.Context, res *models.UserList) error
	UpdateUser(req *models.User, res *models.User) error
	DeleteUser(req *models.GetUserRequest, res *models.GetUserRequest) error
}

type authServiceClient struct {
	client *rpc.Client
}

func NewAuthServiceClient(client *rpc.Client) AuthSvcClient {
	return &authServiceClient{client: client}
}

func (asc *authServiceClient) SignUp(req *models.SignUpRequest, res *models.User) error {
	return asc.client.Call("AuthService.SignUp", req, res)
}

func (asc *authServiceClient) SignIn(req *models.SignInRequest, res *models.User) error {
	return asc.client.Call("AuthService.SignIn", req, res)
}

func (asc *authServiceClient) GetUser(req *models.GetUserRequest, res *models.User) error {
	return asc.client.Call("AuthService.GetUser", req, res)
}

func (asc *authServiceClient) ListUsers(req context.Context, res *models.UserList) error {
	return asc.client.Call("AuthService.ListUsers", req, res)
}

func (asc *authServiceClient) UpdateUser(req *models.User, res *models.User) error {
	return asc.client.Call("AuthService.UpdateUser", req, res)
}

func (asc *authServiceClient) DeleteUser(req *models.GetUserRequest, res *models.GetUserRequest) error {
	return asc.client.Call("AuthService.DeleteUser", req, res)
}
