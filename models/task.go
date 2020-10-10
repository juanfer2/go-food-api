package models

import (
	"gorm.io/gorm"
)

// Types
type Task struct {
	gorm.Model
	ID      int    `json:"ID"`
	Name    string `json:"Name"`
	Content string `json:"Content"`
}

func (t Task) IsNull() bool {
	return t.ID == 0
}
