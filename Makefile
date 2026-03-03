BINARY_NAME=gendiff

.PHONY: build

lint:
	golangci-lint run

build: lint
	go build -o bin/${BINARY_NAME} cmd/${BINARY_NAME}/main.go

run:
	./bin/${BINARY_NAME}