APP_NAME=service-user
MAIN_FILE=main.go
BUILD_DIR=out

.PHONY: all build run test clean tidy

# Default target
all: build

# Build the Go binary
build:
	@echo ">> Building $(APP_NAME)..."
	@go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)

# Run the application
run:
	@echo ">> Running $(APP_NAME)..."
	@go run $(MAIN_FILE)

# Run tests
test:
	@echo ">> Running tests..."
	@go test ./... -v

# Tidy up modules
tidy:
	@echo ">> Running go mod tidy..."
	@go mod tidy

# Clean build artifacts
clean:
	@echo ">> Cleaning up..."
	@rm -rf $(BUILD_DIR)