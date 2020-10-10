package models

import (
	"gorm.io/gorm"
)

// Types
type Food struct {
	gorm.Model
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	ImageURL    string  `json:"image_url"`
	Price       float64 `json:"price"`
	Discount    int     `json:"discount"`

	TypeFoodID int      `json:"type_food_id"`
	TypeFood   TypeFood `gorm:"foreignKey:TypeFoodID"`
}

func (f Food) IsNull() bool {
	return f.ID == 0
}
