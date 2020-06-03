# Makefile to build go-sdk-template

all: build test lint tidy

build:
	go build ./...

test:
	go test `go list ./... | grep -v samples` -tags=integration

lint:
	golangci-lint run

tidy:
	go mod tidy
