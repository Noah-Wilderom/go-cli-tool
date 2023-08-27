# Define the name of the Go source file (change this to your actual filename)
GO_SOURCE_FILE := src/main.go

# Define the name of the output executable
OUTPUT_EXECUTABLE := go-cli

# Define the installation directory for the executable
INSTALL_DIR_UNIX := /usr/bin
INSTALL_DIR_WIN := C:\Program Files\GoCli

# Determine the user's operating system
ifeq ($(OS),Windows_NT)
	OS_TYPE := Windows
	GO_COMPILER := go
else
	UNAME := $(shell uname)
	ifeq ($(UNAME), Darwin)  # macOS
		OS_TYPE := macOS
		GO_COMPILER := go
	else ifeq ($(UNAME), Linux)
		OS_TYPE := Linux
		GO_COMPILER := go
	else
		@echo "Unsupported operating system"
		@exit 1
	endif
endif

# Define the installation directory based on the OS
ifeq ($(OS_TYPE),Windows)
	INSTALL_DIR := $(INSTALL_DIR_WIN)
else
	INSTALL_DIR := $(INSTALL_DIR_UNIX)
endif

# Build and install targets
all: build install

build:
	$(GO_COMPILER) build -o $(OUTPUT_EXECUTABLE) $(GO_SOURCE_FILE)

install:
ifeq ($(OS_TYPE),Windows)
	copy $(OUTPUT_EXECUTABLE) "$(INSTALL_DIR)\$(OUTPUT_EXECUTABLE)"
else
	sudo mv $(OUTPUT_EXECUTABLE) $(INSTALL_DIR)/$(OUTPUT_EXECUTABLE)
endif

.PHONY: all build install