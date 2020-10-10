package models

import (
	"gorm.io/gorm"
)

// Types
type Order struct {
	gorm.Model
	ID          int     `json:"id"`
	Observation string  `json:"observation"`
	Total       float64 `json:"total"`

	UserID int  `json:"user_id"`
	User   User `gorm:"foreignKey:UserID"`
}

func (f Order) IsNull() bool {
	return f.ID == 0
}
