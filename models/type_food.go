package models

import (
	"gorm.io/gorm"
)

// Types
type TypeFood struct {
	gorm.Model
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (tf TypeFood) IsNull() bool {
	return tf.ID == 0
}
