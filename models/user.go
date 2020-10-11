package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        int    `json:"ID"`
	Username  string `json:"Username" gorm:"unique;not null"`
	Password  string `json:"Password"`
	Token     string `json:"Token"`
	SecretKey string `json:"SecretKey"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u User) IsNull() bool {
	return u.ID == 0
}
