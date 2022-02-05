// Package authclient provides a client to the auth service
package authclient

import (
	"net/rpc"

	"github.com/rs401/lg/auth/models"
)

// AuthSvcClient interface defining client methods
type AuthSvcClient interface {
	SignUp(req *models.SignUpRequest, res *models.User) error
	SignIn(req *models.SignInRequest, res *models.User) error
	GetUser(req *models.GetUserRequest, res *models.User) error
	ListUsers(req string, res *models.GetUsersResponse) error
	UpdateUser(req *models.User, res *models.User) error
	DeleteUser(req *models.GetUserRequest, res *models.GetUserRequest) error
}

type authServiceClient struct {
	client *rpc.Client
}

// NewAuthServiceClient takes a pointer to an rpc.Client and returns an
// AuthSvcClient.
func NewAuthServiceClient(client *rpc.Client) AuthSvcClient {
	return &authServiceClient{client: client}
}

// SignUp calls the AuthService.SignUp method
func (asc *authServiceClient) SignUp(req *models.SignUpRequest, res *models.User) error {
	return asc.client.Call("AuthService.SignUp", req, res)
}

// SignIn calls the AuthService.SignIn method
func (asc *authServiceClient) SignIn(req *models.SignInRequest, res *models.User) error {
	return asc.client.Call("AuthService.SignIn", req, res)
}

// GetUser calls the AuthService.GetUser method
func (asc *authServiceClient) GetUser(req *models.GetUserRequest, res *models.User) error {
	return asc.client.Call("AuthService.GetUser", req, res)
}

// ListUsers calls the AuthService.ListUsers method
func (asc *authServiceClient) ListUsers(req string, res *models.GetUsersResponse) error {
	return asc.client.Call("AuthService.ListUsers", req, res)
}

// UpdateUser calls the AuthService.UpdateUser method
func (asc *authServiceClient) UpdateUser(req *models.User, res *models.User) error {
	return asc.client.Call("AuthService.UpdateUser", req, res)
}

// DeleteUser calls the AuthService.DeleteUser method
func (asc *authServiceClient) DeleteUser(req *models.GetUserRequest, res *models.GetUserRequest) error {
	return asc.client.Call("AuthService.DeleteUser", req, res)
}
