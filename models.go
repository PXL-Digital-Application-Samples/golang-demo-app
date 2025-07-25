package main

import "sync"

// User represents a user in the system
type User struct {
	ID    int    `json:"id" example:"1"`
	Name  string `json:"name" example:"Alice" binding:"required"`
	Email string `json:"email" example:"alice@example.com" binding:"required"`
}

// CreateUserRequest represents the request body for creating a user
type CreateUserRequest struct {
	Name  string `json:"name" example:"Alice" binding:"required"`
	Email string `json:"email" example:"alice@example.com" binding:"required"`
}

// UpdateUserRequest represents the request body for updating a user
type UpdateUserRequest struct {
	Name  string `json:"name,omitempty" example:"Alice Updated"`
	Email string `json:"email,omitempty" example:"alice.updated@example.com"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error" example:"User not found"`
}

// MessageResponse represents a success message response
type MessageResponse struct {
	Message string `json:"message" example:"User deleted"`
}

// UserStore manages the in-memory user storage
type UserStore struct {
	users  map[int]*User
	nextID int
	mu     sync.RWMutex
}

// Global user store instance
var userStore *UserStore

// InitializeUsers creates and seeds the user store
func InitializeUsers() {
	userStore = &UserStore{
		users:  make(map[int]*User),
		nextID: 1,
	}
	
	// Seed initial users
	seedUsers := []CreateUserRequest{
		{Name: "Alice", Email: "alice@example.com"},
		{Name: "Bob", Email: "bob@example.com"},
		{Name: "Charlie", Email: "charlie@example.com"},
	}
	
	for _, u := range seedUsers {
		userStore.CreateUser(u.Name, u.Email)
	}
}

// CreateUser creates a new user in the store
func (s *UserStore) CreateUser(name, email string) *User {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	user := &User{
		ID:    s.nextID,
		Name:  name,
		Email: email,
	}
	s.users[s.nextID] = user
	s.nextID++
	
	return user
}

// GetAllUsers returns all users from the store
func (s *UserStore) GetAllUsers() []User {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	users := make([]User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, *user)
	}
	return users
}

// GetUser returns a user by ID
func (s *UserStore) GetUser(id int) (*User, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	user, exists := s.users[id]
	if !exists {
		return nil, false
	}
	return user, true
}

// UpdateUser updates an existing user
func (s *UserStore) UpdateUser(id int, name, email string) (*User, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	user, exists := s.users[id]
	if !exists {
		return nil, false
	}
	
	if name != "" {
		user.Name = name
	}
	if email != "" {
		user.Email = email
	}
	
	return user, true
}

// DeleteUser removes a user from the store
func (s *UserStore) DeleteUser(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	_, exists := s.users[id]
	if !exists {
		return false
	}
	
	delete(s.users, id)
	return true
}

// ResetForTesting clears the store (used in tests)
func (s *UserStore) ResetForTesting() {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	s.users = make(map[int]*User)
	s.nextID = 1
}