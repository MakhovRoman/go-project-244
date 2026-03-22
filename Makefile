BINARY_NAME=gendiff

.PHONY: build lint run

lint:
	golangci-lint run

test:
	go test ./...

build: lint test
	go build -o bin/${BINARY_NAME} cmd/${BINARY_NAME}/main.go

run:
	./bin/${BINARY_NAME}

test-build:
	go build -o bin/${BINARY_NAME} cmd/${BINARY_NAME}/main.go