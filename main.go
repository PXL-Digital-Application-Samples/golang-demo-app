package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "user-management-api/docs"
)

// @title User Management API
// @version 1.0
// @description A lightweight Go Gin-based REST API for managing users (Create, Read, Update, Delete) with in-memory data storage.
// @host localhost:5000
// @BasePath /
func main() {
	// Initialize the Gin router
	router := gin.Default()

	// Initialize the user store
	InitializeUsers()

	// Setup routes
	SetupRoutes(router)

	// Setup Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/swagger/index.html")
	})

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	// Start the server
	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}