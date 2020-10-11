package models

import (
	"time"

	"gorm.io/gorm"
)

// Types
type Food struct {
	gorm.Model
	ID          uint    `json:"id" gorm:"primary_key"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	ImageURL    string  `json:"image_url"`
	Price       float64 `json:"price"`
	Discount    int     `json:"discount"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	TypeFoodID int      `json:"type_food_id"`
	TypeFood   TypeFood `json:"type_food" gorm:"foreignKey:TypeFoodID"`
}

func (f Food) IsNull() bool {
	return f.ID == 0
}
