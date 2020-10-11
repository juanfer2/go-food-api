package models

import (
	"time"

	"gorm.io/gorm"
)

// Types
type OrderFood struct {
	gorm.Model
	ID int `json:"id"`

	TotalPrice float64 `json:"total_price"`

	FoodID int  `json:"food_id"`
	Food   Food `gorm:"foreignKey:FoodID"`

	OrderID int   `json:"order_id"`
	Order   Order `gorm:"foreignKey:OrderID"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (f OrderFood) IsNull() bool {
	return f.ID == 0
}
