package models

type LoginUserResponse struct {
	ID    uint   `json:"id"`
	Token string `json:"token"`
}
