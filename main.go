package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/shobanchiddarth/go-jwt/controllers"
)


func init() {

}

func main() {
	router := gin.Default()

	router.POST("/signup", controllers.Signup) 

	router.Run(":"+os.Getenv("PORT"))
}
