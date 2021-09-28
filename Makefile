all: fmt build test

build:
	go build ./...
	go mod tidy

lint:
	golangci-lint run

fmt:
	golangci-lint run --fix

.PHONY: test
test:
	go test -v ./...
