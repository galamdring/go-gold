# Makefile for building and running Go-Gold application

# Variables
BINARY_NAME=go-gold

# Build the application
build:
	go build -o $(BINARY_NAME) .

# Run the application
run: build
	./$(BINARY_NAME)

# Clean up build artifacts
clean:
	rm -f $(BINARY_NAME)

# Run tests
test:
	go test ./...

# Lint the code
lint:
	golangci-lint run