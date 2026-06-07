package main

import (
	"github.com/shobanchiddarth/go-jwt/initializers"
	"github.com/shobanchiddarth/go-jwt/models"
)


func init() {
	initializers.ConnectToDB()
}


func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
