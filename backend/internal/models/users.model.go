package models

import (
	"okedigital_user_app/backend/pkg/common"

	"gorm.io/gorm"
)

type Users struct {
	common.Model
	Username  string `json:"username" gorm:"unique"`
	Password  string `json:"password"`
	Email     string `json:"email" gorm:"unique"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (u *Users) TableName() string {
	return "users"
}

func (u *Users) BeforeCreate(db *gorm.DB) (err error) {
	u.ID = u.SetUUID("usr")
	return nil
}

type UserWhere struct {
	ID       string
	Username string
}
