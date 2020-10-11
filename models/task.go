package models

import (
	"time"

	"gorm.io/gorm"
)

// Types
type Task struct {
	gorm.Model
	ID      int    `json:"ID"`
	Name    string `json:"Name"`
	Content string `json:"Content"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (t Task) IsNull() bool {
	return t.ID == 0
}
