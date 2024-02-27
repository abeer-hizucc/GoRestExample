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
	r.POST("/createBook", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PUT("/books/:id", controllers.UpdateBook)
	r.Run()
}
