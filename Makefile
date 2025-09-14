# Makefile for ttrk

# Build variables
BINARY_NAME=ttrk
BUILD_DIR=bin
MAIN_PATH=./cmd/ttrk
PKG_LIST=$(shell go list ./...)

# Version and build information
VERSION?=$(shell git describe --tags --always --dirty)
BUILD_DATE=$(shell date +%Y-%m-%d:%H:%M:%S)
GIT_COMMIT=$(shell git rev-parse HEAD)
LDFLAGS=-X github.com/MBH999/ttrk/pkg/version.GitCommit=$(GIT_COMMIT) -X github.com/MBH999/ttrk/pkg/version.BuildDate=$(BUILD_DATE)

.PHONY: all build clean test lint fmt vet help install

all: clean fmt vet test build ## Run all checks and build

build: ## Build the binary
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)

clean: ## Clean build artifacts
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	go clean

test: ## Run tests
	@echo "Running tests..."
	go test -v ./...

lint: ## Run linter
	@echo "Running linter..."
	golangci-lint run

fmt: ## Format code
	@echo "Formatting code..."
	go fmt ./...

vet: ## Run go vet
	@echo "Running go vet..."
	go vet ./...

install: build ## Install binary to GOPATH/bin
	@echo "Installing $(BINARY_NAME)..."
	cp $(BUILD_DIR)/$(BINARY_NAME) $(GOPATH)/bin/

deps: ## Download dependencies
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy

run: build ## Build and run the application
	@echo "Running $(BINARY_NAME)..."
	./$(BUILD_DIR)/$(BINARY_NAME)

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)