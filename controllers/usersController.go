package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour*720).Unix(), // 24*30=720 | expire in 1 month
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if (err!=nil) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create token",
		})
		return
	}


	// send it back, as cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 2592000, "", "", false, false ) // 3600*24*30 | change last 2 to true in production and set domain
	c.JSON(200, gin.H{})
}

