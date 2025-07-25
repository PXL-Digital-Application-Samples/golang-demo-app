.PHONY: all build run test clean swagger deps

# Build the application
build:
	go build -o bin/user-management-api .

# Run the application
run:
	go run .

# Run tests
test:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Clean build artifacts
clean:
	rm -rf bin/
	rm -f coverage.out coverage.html

# Generate swagger documentation
swagger:
	swag init --parseDependency --parseInternal

# Install dependencies
deps:
	go mod download
	go install github.com/swaggo/swag/cmd/swag@latest

# Format code
fmt:
	go fmt ./...

# Run linter
lint:
	golangci-lint run

# Run all checks before committing
check: fmt test

# Initial setup
setup: deps swagger

# Development mode with auto-reload (requires air)
dev:
	air

# Docker build
docker-build:
	docker build -t user-management-api .

# Docker run
docker-run:
	docker run -p 5000:5000 user-management-api