package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

	// create the user

	// respond
}