BINARY_NAME=gendiff

.PHONY: build

build:
	go build -o bin/${BINARY_NAME} cmd/${BINARY_NAME}/main.go

run:
	./bin/${BINARY_NAME}