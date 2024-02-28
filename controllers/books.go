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
func CreateBook(c *gin.Context) {
	//Validate input
	var input models.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book := models.Book{
		Title:  input.Title,
		Author: input.Author,
	}
	models.DB.Create(&book)
	c.JSON(http.StatusOK, gin.H{"data": "Book Created Successfully"})
}
func FindBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}
func UpdateBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var input models.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Model(&book).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": "Book has been successfully updated"})
}
func DeleteBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"data": "Book has been deleted"})
}
