package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/shobanchiddarth/go-jwt/controllers"
	"github.com/shobanchiddarth/go-jwt/initializers"
	"github.com/shobanchiddarth/go-jwt/middleware"
)


func init() {
	initializers.ConnectToDB()
}


func main() {
	router := gin.Default()

	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.GET("/validate", middleware.RequireAuth, controllers.Validate)

	router.Run(":"+os.Getenv("PORT"))
}
