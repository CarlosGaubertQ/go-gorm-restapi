package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string  `gorm:"not_null" json:"firstname"`
	LastName  string  `gorm:"not_null" json:"lastname"`
	Email     string  `gorm:"not_null;unique_index" json:"email"`
	Tasks     []Tasks `json:"tasks"`
}
