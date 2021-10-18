BINARY=engine

all: build

dependencies:
	go mod tidy

build: dependencies
	GOOS=linux GOARCH=amd64 go build -o efishery app/main.go

run:
	go run app/main.go

.PHONY: dependencies build test cover run
