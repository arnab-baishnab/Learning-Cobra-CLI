# Name of the compiled binary
BINARY_NAME := todo

# Directory where the binary will be placed
BUILD_DIR := bin

.PHONY: build run clean help

# Build the Go binary
build: ## Build the Go binary
	@echo "Building the binary..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) ./CobraCLI/main.go

# Run the compiled binary
run: build ## Build and then run the Go binary
	@echo "Running the binary..."
	@./$(BUILD_DIR)/$(BINARY_NAME)

# Clean up the build artifacts
clean: ## Clean up the build artifacts
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)

# Display help for each target
help: ## Show this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'
