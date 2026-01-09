package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	// The parameter here unsures that all the values in this column are unique i.e the Email column is unique
	Email    string `gorm:"unique"`
	Password string
}
