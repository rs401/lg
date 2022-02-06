// Package models provides data structures
package models

// GetUserRequest holds an Id
type GetUserRequest struct {
	Id uint `json:"id"`
}

// SignInRequest holds an Email and Password
type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SignInResponse holds a token
type SignInResponse struct {
	Token string `json:"token"`
}

// SignUpRequest holds a Name, Email and Password
type SignUpRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// GetUsersResponse holds a slice of *Users
type GetUsersResponse struct {
	Users []*User
}
