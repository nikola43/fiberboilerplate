package models

type UserTokenClaims struct {
	ID    uint
	Email string
	Exp   uint
}
