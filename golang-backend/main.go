package main

import (
	"fmt"

	"github.com/betawars/GoLangMessenger/golang-backend/controllers"
	"github.com/betawars/GoLangMessenger/golang-backend/initializers"
	"github.com/betawars/GoLangMessenger/golang-backend/middleware"
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

	// ///////////////////////////////////////////////////Code for CRUD operations//////////////////////////////////////////////////////
	// The table used over here is the "posts" table
	// This endpoint is to create a post in the database
	r.POST("/createPost", controllers.CreatePosts)
	// This endpoint is to get all the data in the database
	r.GET("/getPosts", controllers.GetPosts)
	// This enpoint is to get data with a particular id from the database
	r.GET("/getPost/:id", controllers.GetPost)
	// This endpoint is to update the existing posts from the database
	r.PUT("/updatePost/:id", controllers.UpdatePost)
	// This endpoint is to delete a record from the database
	r.DELETE("deletePost/:id", controllers.DeletePost)

	///////////////////////////////////////////////////////Code for JWT///////////////////////////////////////////////////////////////
	// The table used over here is the "users" table
	// This endpoint is to create new User in the Users table
	r.POST("/signUp", controllers.SignUp)
	// This endpoint is to check if the user exists in the Users table
	r.POST("/login", controllers.Login)
	// TODO Doc
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
}
