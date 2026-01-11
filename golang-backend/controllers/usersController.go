package controllers

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/betawars/GoLangMessenger/golang-backend/initializers"
	"github.com/betawars/GoLangMessenger/golang-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

func Login(c *gin.Context) {

	//  Get email and pass from the req body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
	}

	// Look up for the requested user
	var user models.User
	initializers.DB.Where("email = ?", body.Email).First(&user)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Email or Password",
		})
		return
	}

	// Compare the user with pass

	storedPassword := user.Password
	bodyPassword := body.Password

	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(bodyPassword))

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"information": "Incorrect Password!",
			"details":     err.Error(),
		})
	}

	// Generate the JWT Token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"subject": user.ID,
		"expire":  time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed to create token",
			"errInfo": err.Error(),
		})
		return
	}

	// Send it back
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
