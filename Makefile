.PHONY: help test test-coverage test-race test-short lint fmt vet build clean dev run migrate-up migrate-down migrate-create docker-up docker-down docker-logs docker-test docker-clean

# Default target
help:
	@echo "Available targets:"
	@echo "  make test          - Run all tests"
	@echo "  make test-coverage - Run tests with coverage report"
	@echo "  make test-race     - Run tests with race detector"
	@echo "  make test-short    - Run short tests only"
	@echo "  make lint          - Run linters"
	@echo "  make fmt           - Format code"
	@echo "  make vet           - Run go vet"
	@echo ""
	@echo "Docker targets:"
	@echo "  make docker-up     - Start Booklore server with Docker Compose"
	@echo "  make docker-down   - Stop Docker Compose services"
	@echo "  make docker-logs   - View Docker Compose logs"
	@echo "  make docker-test   - Run tests against Docker-based Booklore server"


# Testing targets
test:
	@echo "Running tests..."
	go test -v ./...

test-coverage:
	@echo "Running tests with coverage..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

test-race:
	@echo "Running tests with race detector..."
	go test -v -race ./...

test-short:
	@echo "Running short tests..."
	go test -v -short ./...

# Run benchmarks
benchmark:
	@echo "Running benchmarks..."
	go test -bench=. -benchmem -run=^$$ ./...

# Code quality targets
lint:
	@echo "Running linters..."
	docker run -t --rm -v ./:/app -w /app golangci/golangci-lint:v2.8.0 golangci-lint run

fmt:
	@echo "Formatting code..."
	go fmt ./...
	gofmt -s -w .

vet:
	@echo "Running go vet..."
	go vet ./...

# Docker targets
docker-up:
	@echo "Starting Booklore server with Docker Compose..."
	docker-compose up -d
	@echo ""
	@echo "Booklore server is starting up. Waiting for health check..."
	@sleep 5
	@echo ""
	@echo "Services running:"
	@echo "  - PostgreSQL: localhost:5432"
	@echo "  - Booklore API: http://localhost:8080"
	@echo ""
	@echo "To run tests against Docker server, use:"
	@echo "  make docker-test"

docker-down:
	@echo "Stopping Docker Compose services..."
	docker-compose down

docker-logs:
	@echo "Showing Docker Compose logs (last 50 lines)..."
	docker-compose logs --tail=50 -f

docker-test:
	@echo "Running tests against Docker-based Booklore server..."
	@echo "Loading Docker environment (.env.docker)..."
	docker-compose exec -T booklore bash -c "echo waiting for server..." || true
	@env $$(cat .env.docker | grep -v '^#' | xargs) go test -v ./pkg/booklore/...

docker-clean:
	@echo "Cleaning up Docker containers and volumes..."
	docker-compose down -v

