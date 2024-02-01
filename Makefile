BUILD_DIR=dist
SERVER_NAME=example-server
CLIENT_NAME=example

.PHONY: all build test docker run tidy clean

all: build test

build:
	mkdir -p ${BUILD_DIR}
	go build -o ${BUILD_DIR}/${SERVER_NAME} -v ./cmd/server/main.go
	go build -o ${BUILD_DIR}/${CLIENT_NAME} -v ./cmd/client/main.go

test:
	go test -v ./...

docker:
	docker build -t mukkak/example-service:latest .

run:
	go run ./cmd/server/main.go

tidy:
	go mod tidy

clean:
	go clean
	rm -rf ${BUILD_DIR}
