package main

import (
	"GoProject/controllers"
	"GoProject/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()
	r.GET("/books", controllers.FindBooks)
	r.Run()
}
