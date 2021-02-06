# TODO: implement when we have any tag at all
#VERSION := $(shell git describe)

VERSION := "v0.0.0"
BUILD := $(shell git rev-parse --short HEAD)
BINARY_NAME="service"
LDFLAGS=-ldflags "-X=$(BINARY_NAME).Version=$(VERSION) -X=$(BINARY_NAME).Build=$(BUILD)"
WORKING_DIR := $(shell pwd)
HOME_DIR := $(HOME)
GOPATH := $(GOPATH)

FORCE:

dev:get up log

clean: down

down:
	docker-compose down --remove-orphans

get:
	go get -t ./...
	
lint:
	docker run --rm -v "$(GOPATH)/pkg":/go/pkg -v $(PWD):/app -w /app golangci/golangci-lint:v1.33.0 golangci-lint run --timeout 5m

log:
	docker-compose logs -f

test:
	go test -cover $(shell go list ./... | grep -v /integration)

up:
	docker-compose up -d