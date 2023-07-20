package controllers

import (
	"bookstore/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllBook(ctx *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)
	ctx.JSON(http.StatusOK, gin.H{
		"successful": true,
		"message":    "fetch books successfully",
		"data":       books,
	})
}

func CreateBook(ctx *gin.Context) {
	var BookBody struct {
		Title  string `json:"title" binding:"required"`
		Author string `json:"author" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&BookBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
		return
	}
	book := models.Book{Title: BookBody.Title, Author: BookBody.Author}
	models.DB.Create(&book)
	ctx.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func GetSingleBook(ctx *gin.Context) {
	// The id from params
	id := ctx.Param("id")
	var book models.Book
	// Check if book with ID exists
	err := models.DB.Where("id = ?", id).First(&book).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return
	}
	// Return if book is found
	ctx.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func UpdateBook(ctx *gin.Context) {
	// Schema for validating input
	var BookBody struct {
		Title  string `json:"title"`
		Author string `json:"author"`
	}
	// get Id from params
	id := ctx.Param("id")
	var book models.Book
	err := models.DB.Where("id = ?", id).First(&book).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	if err := ctx.ShouldBindJSON(&BookBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Model(&book).Updates(BookBody)
	ctx.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")
	var book models.Book
	if err := models.DB.Where("id = ?", id).First(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Record not found!",
		})
		return
	}
	models.DB.Delete(&book)
	ctx.JSON(http.StatusOK, gin.H{
		"mesaage": "Book deleted successfully!",
	})
}
