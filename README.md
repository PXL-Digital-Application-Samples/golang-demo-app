# 🧩 Go User Management API

A lightweight Go Gin-based REST API for managing users (Create, Read, Update, Delete) with in-memory data storage. Includes auto-seeded user data and interactive Swagger (OpenAPI) documentation served at the root (`/`).

---

## 📦 Features

- ✅ In-memory storage with thread-safe operations
- 🔄 Full CRUD endpoints: `POST`, `GET`, `PUT`, `DELETE`
- 🧪 Swagger UI for testing & docs at `/` 
- 🚀 High performance with Go's concurrency
- 🧰 Easily extendable to use persistent databases (e.g., PostgreSQL, MongoDB)
- 🔒 Thread-safe operations with mutex locks

---

## 🚀 Getting Started

### Requirements

- Go 1.21+
- Make (optional, for using Makefile commands)

### Installation

```bash
# Clone the repository
git clone <repository-url>
cd user-management-api

# Download dependencies
go mod download

# Install swag CLI tool for generating Swagger docs
go install github.com/swaggo/swag/cmd/swag@latest

# If swag is not in PATH after installation, add Go bin to PATH:
# export PATH=$PATH:$(go env GOPATH)/bin

# Generate Swagger documentation
swag init --parseDependency --parseInternal

# Run tests to verify setup
go test -v ./...
```

### Quick Start with Make

```bash
# Initial setup (install deps + generate swagger)
make setup

# Run the application
make run

# Run tests
make test

# Build binary
make build
```

### Manual Commands

```bash
# Generate Swagger docs
swag init --parseDependency --parseInternal

# Run the application
go run .

# Build the application
go build -o bin/user-management-api .

# Run tests
go test -v ./...

# Run tests with coverage
go test -v -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```

---

## 📋 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET    | `/` | Swagger UI Documentation |
| POST   | `/users` | Create a new user |
| GET    | `/users` | Get all users |
| GET    | `/users/:id` | Get a specific user |
| PUT    | `/users/:id` | Update a user |
| DELETE | `/users/:id` | Delete a user |

---

## 🧪 Testing
**
```bash
# Run all tests
make test

# Run tests with verbose output
go test -v ./...

# Run tests with coverage report
make test-coverage
```

---

## 🛠️ Development

### Project Structure

```
.
├── main.go           # Application entry point
├── models.go         # Data models and storage
├── handlers.go       # HTTP request handlers
├── handlers_test.go  # Unit tests
├── go.mod           # Go module file
├── go.sum           # Go module checksums
├── Makefile         # Build automation
├── docs/            # Auto-generated Swagger docs
└── bin/             # Compiled binaries
```

### Environment Variables

- `PORT` - Server port (default: 5000)

### Advanced Development Setup

#### Hot-Reload Development
For development with automatic reloading when files change:
```bash
# Install air for hot-reload
go install github.com/cosmtrek/air@latest

# Run with hot-reload
air
```

#### Environment Configuration
If you have environment-specific configurations:
```bash
# Copy example environment file (if available)
cp .env.example .env
```

### Example Usage

#### Create a User
```bash
curl -X POST http://localhost:5000/users \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "email": "john@example.com"}'
```

#### Get All Users
```bash
curl http://localhost:5000/users
```

#### Get a Specific User
```bash
curl http://localhost:5000/users/1
```

#### Update a User
```bash
curl -X PUT http://localhost:5000/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "John Updated"}'
```

#### Delete a User
```bash
curl -X DELETE http://localhost:5000/users/1
```
