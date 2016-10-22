.PHONY: build test deps compile

export GOPATH := ${PWD}/vendor:${PWD}
export GOBIN := ${PWD}/vendor/bin

NAME := solution-a

default: build

deps:
	@go get github.com/go-sql-driver/mysql
	@go get github.com/icrowley/fake
	@go get gopkg.in/yaml.v2

DP := /usr/src/solution
build:  deps
	@echo Building...
	@rm -rf ./bin/*
	go build -v -o ./bin/$(NAME) ./src/main.go
	@echo Done.

compile:
	@# TODO GOOS, GOARCH - should to be in config
	@docker run --rm -v $(PWD):$(DP) -w $(DP) -e GOPATH=$(DP) -e GOOS="darwin" -e GOARCH="amd64" golang:1.6 sh -c 'make build'

test:
	@go test test/*.go

clean-deps:
	rm -rf ./vendor/*
	rm -rf ./pkg/*
