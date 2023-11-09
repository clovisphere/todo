.PHONY: build-docker-image cover run-docker-image start test test-integration

# Latest commit hash
RELEASE=$(shell git rev-parse --short=8 HEAD)

build-docker-image:
	docker build -t todo-api --build-arg RELEASE=$(RELEASE) .

cover:
	go tool cover -html=cover.out

run-docker-image:
	docker run -p 6969:9000 -e HOST="" todo-api

start:
	go run cmd/server/*.go

test:
	go test -coverprofile=cover.out -short ./...

test-integration:
	go test -coverprofile=cover.out -p 1 ./...
