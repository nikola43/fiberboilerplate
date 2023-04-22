package services

import (
	"github.com/nikola43/fiberboilerplate/internal/api/v1/database"
	dbmodels "github.com/nikola43/fiberboilerplate/internal/api/v1/models/db"
)

// GetUserById returns a user by id
func GetUserById(id uint64) (*dbmodels.User, error) {
	user := new(dbmodels.User)
	user.ID = uint(id)

	err := database.GormDB.First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Update updates a user
func Update(user *dbmodels.User) (*dbmodels.User, error) {
	// check if user exists
	err := database.GormDB.First(&user).Error
	if err != nil {
		return nil, err
	}

	tx := database.GormDB.Save(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

// Delete deletes a user
func DeleteUser(id uint64) error {
	user := new(dbmodels.User)
	user.ID = uint(id)

	// check if user exists
	err := database.GormDB.First(&user).Error
	if err != nil {
		return err
	}

	tx := database.GormDB.Delete(&user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
