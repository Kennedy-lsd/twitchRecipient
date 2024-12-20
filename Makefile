APP_NAME := StreamAlert
BUILD_DIR := bin
VERSION := 1.0.0

build: $(BUILD_DIR)
	cp .env $(BUILD_DIR)/.env
	go build -o $(BUILD_DIR)/$(APP_NAME)

$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

# Build for Windows
build-windows:
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME).exe

clean:
	rm -rf $(BUILD_DIR)

install: $(BUILD_DIR)
	sudo cp $(BUILD_DIR)/$(APP_NAME) /usr/local/bin/$(APP_NAME)
	@echo "Program installed globally. You can run '$(APP_NAME)' from anywhere."

help:
	@echo "Makefile commands:"
	@echo "  build         Build the binary for the current platform"
	@echo "  build-windows Build the binary for Windows"
	@echo "  install       Install the binary globally"
	@echo "  clean         Remove build artifacts"
	@echo "  help          Show this help message"
