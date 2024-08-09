# Makefile for building and running Go-Gold application

# Variables
BINARY_NAME=build/go-gold

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
	golangci-lint -v run

check: test lint

# Fix lint issues and format code
fix:
	golangci-lint run --fix
	go fmt ./...

live:
	air --build.cmd "go build -o $(BINARY_NAME) ." --build.bin "./${BINARY_NAME}" --build.exclude_dir "templates,build,examples"