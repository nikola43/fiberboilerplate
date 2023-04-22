package models

type SignupUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
