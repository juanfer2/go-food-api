package models

import (
	"time"

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

	OrderFoods []OrderFood `json:"order_foods" gorm:"foreignKey:OrderID"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (f Order) IsNull() bool {
	return f.ID == 0
}
