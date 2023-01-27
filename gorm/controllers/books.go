package controllers

import (
	"net/http"

	"github.com/fostemi/golang-playground/example/gorm"
	"github.com/gin-gonic/gin"
)

// GET /books
func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}
