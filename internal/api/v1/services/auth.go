package services

import (
	"errors"

	"github.com/nikola43/fiberboilerplate/internal/api/v1/auth"
	"github.com/nikola43/fiberboilerplate/internal/api/v1/database"
	dbmodels "github.com/nikola43/fiberboilerplate/internal/api/v1/models/db"

	requestmodels "github.com/nikola43/fiberboilerplate/internal/api/v1/models/requests"
	"github.com/nikola43/fiberboilerplate/pkg/utils"
)

func LoginClient(email, password string) (*requestmodels.LoginUserResponse, error) {
	user := new(dbmodels.User)
	pk := auth.PrivateKey

	err := database.GormDB.
		Where("email = ?", email).
		Find(&user).Error
	if err != nil {
		return nil, err
	}

	if !utils.ComparePasswords(user.Password, []byte(password)) {
		return nil, errors.New("invalid email or password")
	}

	token, err := utils.GenerateClientToken(pk, user.Email, user.ID)
	if err != nil {
		return nil, err
	}

	return &requestmodels.LoginUserResponse{
		ID:    user.ID,
		Token: token,
	}, err
}

func SignUpClient(request *requestmodels.SignupUserRequest) error {
	user := new(dbmodels.User)

	err := database.GormDB.
		Where("email = ?", request.Email).
		Find(&user).Error
	if err != nil {
		return err
	}

	if user.ID > 0 {
		return errors.New("user already exists")
	}

	// create user on database
	user = &dbmodels.User{
		Email:    request.Email,
		Password: utils.HashPassword([]byte(request.Password)),
	}

	result := database.GormDB.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
