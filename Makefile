# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=hamara
BINARY_UNIX=$(BINARY_NAME)_unix
VERSION=$(shell git describe --abbrev=0 --tags)

all: test build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v

.PHONY: test # Run all tests
test:
	$(GOTEST) -v ./...

.PHONY: clean # Clean the project and removes existing binaries
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

.PHONY: run # Run the generated binary for the current platform
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)

.PHONY: deps # Generate list of targets with descriptions
deps:
	$(GOCMD) mod download

.PHONY: build-linux # Compile the binary for the Linux OS
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

.PHONY: version # Print the current tagged version
version:
	@echo $(VERSION)

.PHONY: docker-build # Build a Docker Image
docker-build:
	docker build -t trivago/$(BINARY_NAME):$(VERSION) .

.PHONY: help # Generate list of targets with descriptions
help:
	@echo "Available targets:"
	@grep '^.PHONY: .* #' $(MAKEFILE_LIST) | sed 's/\.PHONY: \(.*\) # \(.*\)/\1:\2/' | column -c2 -t -s ':'

.DEFAULT_GOAL := help
