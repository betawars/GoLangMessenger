// This file is to run when we want to create a DB at our DB server.
// This needs to be done before doing any operation the DB so that the tabel exists.
// use the command "go run migrate/migrate.go" to run this file
package main

import (
	"github.com/betawars/GoLangMessenger/golang-backend/initializers"
	// "github.com/betawars/GoLangMessenger/golang-backend/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	// initializers.DB.AutoMigrate(&models.Post{})
	// initializers.DB.AutoMigrate(&models.User{})
}
