.PHONY: build test deps

export GOPATH := ${PWD}/vendor:${PWD}
export GOBIN := ${PWD}/vendor/bin


NAME := solution-a

default: build

deps:
	@go get github.com/go-sql-driver/mysql

build:  deps
	@echo Building...
	rm -rf ./bin/*
	go build -v -o ./bin/$(NAME) ./src/main.go
	@#docker run --rm -v "$PWD":/usr/src/solution -w /usr/src/solution golang:1.6 make compile
	@echo Done.

compile:
	@go build -v -o bin/solution-a src/main.go

test:
	@go test test/*.go

clean-deps:
	rm -rf ./vendor/*
	rm -rf ./pkg/*
