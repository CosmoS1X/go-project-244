RUN_ARGS ?=
APP_NAME := gendiff

.PHONY: fmt tidy test lint build run vuln clean

.DEFAULT_GOAL := build

fmt:
	golangci-lint fmt

tidy:
	go mod tidy

test:
	go test -v ./...

test-race:
	go test -race -v ./...

test-bench:
	go test -bench=. -benchmem

test-coverage:
	go test -v -cover -coverprofile=coverage.out ./...

show-coverage:
	go tool cover -html=coverage.out

lint: fmt
	golangci-lint run

lint-fix:
	golangci-lint run --fix

build:
	go build -o bin/$(APP_NAME) ./cmd/$(APP_NAME)

run:
	go run cmd/$(APP_NAME)/main.go $(RUN_ARGS)

install:
	go install ./cmd/$(APP_NAME)

vuln:
	govulncheck ./...

clean:
	rm -rf bin/
