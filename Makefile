BUILD_DIR=dist
CLIENT_NAME=example
SERVER_NAME=example-server

.PHONY: all tidy build test run clean

all: build test

tidy:
	go mod tidy

build:
	mkdir -p ${BUILD_DIR}
	go build -o ${BUILD_DIR}/${SERVER_NAME} -v ./cmd/server/main.go
	go build -o ${BUILD_DIR}/${CLIENT_NAME} -v ./cmd/client/main.go

test:
	go test -v ./...

run:
	go run .

clean:
	go clean
	rm -rf ${BUILD_DIR}
