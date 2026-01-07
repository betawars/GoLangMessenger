package main

import (
	"fmt"

	"github.com/betawars/GoLangMessenger/golang-backend/controllers"
	"github.com/betawars/GoLangMessenger/golang-backend/initializers"
	"github.com/gin-gonic/gin"
)

// this function runs before main and is used to setup the server at the desired port(default 8080) from the .env file - For now
func init() {
	fmt.Println("-----------------------------------------------------------START'S HERE-------------------------------------------------------------")
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.CreatePosts)
	r.Run()
}
