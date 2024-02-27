package controllers

import (
	"GoProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
	// Ensure you import your models package
)

// FindBooks retrieves all books from the database and sends them as a JSON response.
func FindBooks(c *gin.Context) {
	var books []models.Book

	// Query the database
	result := models.DB.Find(&books)
	if result.Error != nil {
		// Log the error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve books"})
		return
	}

	// Send the books as a JSON response
	c.JSON(http.StatusOK, gin.H{"data": books})
}
