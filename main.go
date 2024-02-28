package main

import (
	"GoProject/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	dbErr := models.ConnectDatabase()
	if dbErr != nil {
		return
	}
	CRUDRoutes(r)
	rerr := r.Run()
	if rerr != nil {
		return
	}

}
