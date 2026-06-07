package main

import (
	"os"

	"github.com/gin-gonic/gin"
)


func init() {

}

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
		"message": "pong",
		})
	})

	router.Run(":"+os.Getenv("PORT"))
}
