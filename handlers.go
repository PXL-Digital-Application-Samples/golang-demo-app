package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all the API routes
func SetupRoutes(router *gin.Engine) {
	users := router.Group("/users")
	{
		users.POST("", CreateUserHandler)
		users.GET("", GetUsersHandler)
		users.GET("/:id", GetUserHandler)
		users.PUT("/:id", UpdateUserHandler)
		users.DELETE("/:id", DeleteUserHandler)
	}
}

// CreateUserHandler creates a new user
// @Summary Create a new user
// @Description Create a new user with name and email
// @Tags Users
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "User data"
// @Success 201 {object} User
// @Failure 400 {object} ErrorResponse
// @Router /users [post]
func CreateUserHandler(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Name and email are required"})
		return
	}

	user := userStore.CreateUser(req.Name, req.Email)
	c.JSON(http.StatusCreated, user)
}

// GetUsersHandler returns all users
// @Summary Get all users
// @Description Get a list of all users
// @Tags Users
// @Produce json
// @Success 200 {array} User
// @Router /users [get]
func GetUsersHandler(c *gin.Context) {
	users := userStore.GetAllUsers()
	c.JSON(http.StatusOK, users)
}

// GetUserHandler returns a single user by ID
// @Summary Get a user by ID
// @Description Get user data by user ID
// @Tags Users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Failure 404 {object} ErrorResponse
// @Router /users/{id} [get]
func GetUserHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid user ID"})
		return
	}

	user, exists := userStore.GetUser(id)
	if !exists {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUserHandler updates an existing user
// @Summary Update a user
// @Description Update user data by user ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body UpdateUserRequest true "Updated user data"
// @Success 200 {object} User
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /users/{id} [put]
func UpdateUserHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid user ID"})
		return
	}

	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "No data provided"})
		return
	}

	user, exists := userStore.UpdateUser(id, req.Name, req.Email)
	if !exists {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUserHandler deletes a user
// @Summary Delete a user
// @Description Delete a user by user ID
// @Tags Users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} MessageResponse
// @Failure 404 {object} ErrorResponse
// @Router /users/{id} [delete]
func DeleteUserHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid user ID"})
		return
	}

	if !userStore.DeleteUser(id) {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "User not found"})
		return
	}

	c.JSON(http.StatusOK, MessageResponse{Message: "User deleted"})
}