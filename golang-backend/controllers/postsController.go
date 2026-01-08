// Documentation for the CRUD controllers can be found at "grom.io/docs/create.html"
package controllers

import (
	"github.com/betawars/GoLangMessenger/golang-backend/initializers"
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

func GetPosts(c *gin.Context) {
	//Get the posts

	var posts []models.Post
	initializers.DB.Find(&posts)

	//Respond with the recieved posts

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func GetPost(c *gin.Context) {

	// get ID off URL(this comes from the URL for ex: http://localhost:3000/post/:id) here the ":id" is the dynamic parameter which is fetched in the line below
	id := c.Param("id")

	// Get a single post
	var post models.Post
	intializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})

}

func UpdatePost(c *gin.Context) {

	//Get the ID from the URL to be updated

	id := c.Param("id")

	// Find the post we are updating
	var post models.Post
	initializers.DB.First(&post, id)

	// Get the data off the request body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// Respond with the updated post

	c.JSON(200, gin.H{
		"post": post,
	})

}

func DeletePost(c *gin.Context) {
	// Get the ID off the URL
	id := c.Param("id")

	// Get the post using the id
	var post models.Post
	initializers.DB.First(&post, id)

	// Delete the post
	initializers.DB.Where("id = ?", id).Delete(&post)

	// Send the response
	c.JSON(200, gin.H{
		"post": post,
	})
}
