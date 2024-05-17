# Main package name
MAIN_PACKAGE = ./src

APP_NAME = game
BUILD_DIR = ./bin

# Default target
default: build run

# Build target
build:
	go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_PACKAGE)

# Run target
run:
	$(BUILD_DIR)/$(APP_NAME)

# Clean target
clean:
	go clean
	rm -f $(BINARY_NAME)
