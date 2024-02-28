package main

import (
	"GoProject/controllers"
	"github.com/gin-gonic/gin"
)

func CRUDRoutes(r *gin.Engine) {
	r.GET("/books", controllers.FindBooks)
	r.POST("/createBook", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PUT("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)
}
