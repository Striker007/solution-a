.PHONY: build test init

export GOPATH := ${PWD}/vendor:${PWD}
export GOBIN := ${PWD}/vendor/bin


NAME := solution-a

default: build

init:
        go get github.com/go-sql-driver/mysql

build:
	@echo Building...
	go build -v -o ./bin/$(NAME) ./src/*.go
	@echo Done.

test:
	@go test test/*.go
