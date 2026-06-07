package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shobanchiddarth/go-jwt/initializers"
	"github.com/shobanchiddarth/go-jwt/models"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	// get email, password off of request body
	var body struct {
		Email string
		Password string
	}

	if (c.Bind(&body) != nil) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	// hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if (err!=nil) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	// create the user
	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if (result.Error!=nil) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "could not signup",
		})
		return
	}

	// respond
	c.JSON(200, gin.H{
		"message":"signup successful",
	})
}

func Login(c *gin.Context) {
	// get email and password off request body
	var body struct {
		Email string
		Password string
	}

	if (c.Bind(&body) != nil) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	// look up requested user
	var user models.User

	initializers.DB.First(&user, "email = ?", body.Email)

	if (user.ID == 0) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	// compare hashes
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if (err != nil) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid username or password",
		})
		return
	}


	// generate jwt token


	// send it back
}

