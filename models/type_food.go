package models

import (
	"time"

	"gorm.io/gorm"
)

// Types
type TypeFood struct {
	gorm.Model
	ID   int    `json:"id"`
	Name string `json:"name"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (tf TypeFood) IsNull() bool {
	return tf.ID == 0
}
