package main

import (
	"bookstore/controllers"
	"bookstore/initializers"
	"bookstore/models"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	initializers.LoadEnvVariable()
	models.ConnectDatabase()
}

func main() {
	r := gin.Default()
	r.GET("/api/v1/books", controllers.GetAllBook)
	r.POST("/api/v1/books", controllers.CreateBook)
	r.GET("/api/v1/books/:id", controllers.GetSingleBook)
	r.PUT("/api/v1/books/:id", controllers.UpdateBook)
	r.DELETE("/api/v1/books/:id", controllers.DeleteBook)
	err := r.Run()
	if err != nil {
		log.Fatal("Failed to start app!")
	}
}
