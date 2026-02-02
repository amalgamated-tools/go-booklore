.PHONY: help test test-coverage test-race test-short lint fmt vet build clean dev run migrate-up migrate-down migrate-create

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
