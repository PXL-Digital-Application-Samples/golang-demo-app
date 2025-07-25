package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	SetupRoutes(router)
	return router
}

func setupTest() {
	// Initialize a fresh user store for each test
	userStore = &UserStore{
		users:  make(map[int]*User),
		nextID: 1,
	}
	// Add test users
	userStore.users[1] = &User{ID: 1, Name: "Alice", Email: "alice@example.com"}
	userStore.users[2] = &User{ID: 2, Name: "Bob", Email: "bob@example.com"}
	userStore.nextID = 3
}

func TestGetAllUsers(t *testing.T) {
	setupTest()
	router := setupTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var users []User
	err := json.Unmarshal(w.Body.Bytes(), &users)
	assert.NoError(t, err)
	assert.Len(t, users, 2)
}

func TestGetSingleUser(t *testing.T) {
	setupTest()
	router := setupTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var user User
	err := json.Unmarshal(w.Body.Bytes(), &user)
	assert.NoError(t, err)
	assert.Equal(t, "Alice", user.Name)
}

func TestGetUserNotFound(t *testing.T) {
	setupTest()
	router := setupTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/99", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCreateUser(t *testing.T) {
	setupTest()
	router := setupTestRouter()

	newUser := CreateUserRequest{
		Name:  "Charlie",
		Email: "charlie@example.com",
	}
	jsonValue, _ := json.Marshal(newUser)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var user User
	err := json.Unmarshal(w.Body.Bytes(), &user)
	assert.NoError(t, err)
	assert.Equal(t, "Charlie", user.Name)
	assert.Equal(t, "charlie@example.com", user.Email)
	assert.NotZero(t, user.ID)
}

func TestCreateUserInvalid(t *testing.T) {
	setupTest()
	router := setupTestRouter()

	incompleteUser := map[string]string{
		"name": "Incomplete",
	}
	jsonValue, _ := json.Marshal(incompleteUser)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateUser(t *testing.T) {
	setupTest()
	router := setupTestRouter()

	updateData := UpdateUserRequest{
		Name: "Alice Updated",
	}
	jsonValue, _ := json.Marshal(updateData)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/users/1", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var user User
	err := json.Unmarshal(w.Body.Bytes(), &user)
	assert.NoError(t, err)
	assert.Equal(t, "Alice Updated", user.Name)
	assert.Equal(t, "alice@example.com", user.Email) // Email should remain unchanged
}

func TestDeleteUser(t *testing.T) {
	setupTest()
	router := setupTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/users/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// Verify user is deleted
	_, exists := userStore.GetUser(1)
	assert.False(t, exists)
}

func TestDeleteNonexistentUser(t *testing.T) {
	setupTest()
	router := setupTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/users/999", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}