// This file is to run when we want to create a DB at our DB server.
package main

import (
	"github.com/betawars/GoLangMessenger/golang-backend/initializers"
	"github.com/betawars/GoLangMessenger/golang-backend/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
