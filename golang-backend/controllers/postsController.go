// Documentation for the CRUD controllers can be found at "grom.io/docs/create.html"
package controllers

import (
	intializers "github.com/betawars/GoLangMessenger/golang-backend/initializers"
	"github.com/betawars/GoLangMessenger/golang-backend/models"
	"github.com/gin-gonic/gin"
)

func CreatePosts(c *gin.Context) {

	//Get data off req body

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//Create a post

	post := models.Post{Title: body.Title, Body: body.Body}
	result := intializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

// func GetPosts(c *gin.Context) {
// 	//Get the posts

// 	//Respond with the recieved posts
// }
