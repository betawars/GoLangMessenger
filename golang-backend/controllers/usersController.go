package controllers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/betawars/GoLangMessenger/golang-backend/initializers"
	"github.com/betawars/GoLangMessenger/golang-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	// get the email/pass
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Hash the password

	hash, error := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	// Create the user
	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {

		// This is the method of high reliability where postgres error code is checked in this case "23505" for duplicate user as it will stay throught the DBs
		var pgErr *pgconn.PgError

		// This checks if the error is a Postgres error and "casts" it to pgErr
		if errors.As(result.Error, &pgErr) {
			if pgErr.Code == "23505" { // 23505 is the unique_violation code
				c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
				return
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to create user",
			})
		}

		// This is the method in which it is just checked if "duplicate" is present in the response but it is low reliability
		if strings.Contains(result.Error.Error(), "duplicate") {
			c.JSON(http.StatusConflict, gin.H{
				"error": "User already exists!",
			})

		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to create user",
			})
		}

		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"information": "User Created Successfully!",
	})

}
