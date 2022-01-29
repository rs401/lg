package models

type GetUserRequest struct {
	Id uint
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInResponse struct {
	Token string `json:"token"`
}
