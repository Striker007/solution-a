.PHONY: build test

export GOPATH := ${PWD}/vendor:${PWD}
export GOBIN := ${PWD}/vendor/bin


NAME := solution-a

default: build

build:
	@echo Building...
	go build -v -o ./bin/$(NAME) ./src/*.go
	@echo Done.

test:
	@go test test/*.go
