# Define the name of the binary
BINARY_NAME=simplebank

# Default target executed when no arguments are given to make
.PHONY: all
all: build

# Rule to build the Go project
.PHONY: build
build:
	go build -o $(BINARY_NAME) ./cmd/.

# Rule to run the application
.PHONY: run
run:
	go run ./cmd/.

# Rule to clean the project directory
.PHONY: clean
clean:
	rm -f $(BINARY_NAME)
	