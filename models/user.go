package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        int    `json:"ID"`
	Username  string `json:"Username" gorm:"unique;not null"`
	Password  string `json:"Password"`
	Token     string `json:"Token"`
	SecretKey string `json:"SecretKey"`
}

func (u User) IsNull() bool {
	return u.ID == 0
}
