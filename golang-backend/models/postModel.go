//This file is where the model of the database is declared. When we are creating a dable in the DB and we need a model for it, this is where we come.

package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title string
	Body  string
}
