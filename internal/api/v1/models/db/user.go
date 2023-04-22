package models

import "github.com/nikola43/fiberboilerplate/internal/api/v1/models/base"

type User struct {
	base.CustomGormModel
	Email    string `gorm:"index; unique; type:varchar(64) not null" json:"email"`
	Password string `gorm:"type:varchar(256) not null; size:256" json:"password"`
}
