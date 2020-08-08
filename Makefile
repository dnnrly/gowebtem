GO111MODULE=on

CURL_BIN ?= curl
GO_BIN ?= go
GORELEASER_BIN ?= goreleaser

PUBLISH_PARAM?=
GO_MOD_PARAM?=-mod vendor
TMP_DIR?=./tmp

BASE_DIR=$(shell pwd)

NAME=gowebtem

export GO111MODULE=on
export GOPROXY=https://proxy.golang.org
export PATH := $(BASE_DIR)/bin:$(PATH)

.PHONY: install deps clean clean-deps test-deps build-deps deps test acceptance-test ci-test lint release update

install:
	go install -v ./cmd/$(NAME)

build:
	go build -v ./cmd/$(NAME)
	docker build .

clean:
	rm -f $(NAME)

clean-deps:
	rm -rf ./bin
	rm -rf ./tmp

./bin/golangci-lint:
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s v1.28.3

./bin/tparse: ./bin ./tmp
	curl -sfL -o ./tmp/tparse.tar.gz https://github.com/mfridman/tparse/releases/download/v0.7.4/tparse_0.7.4_Linux_x86_64.tar.gz
	tar -xf ./tmp/tparse.tar.gz -C ./bin

test-deps: ./bin/tparse ./bin/golangci-lint
	go get -v ./...
	go mod tidy

./bin:
	mkdir ./bin

./tmp:
	mkdir ./tmp

build-deps:

deps: build-deps test-deps

test: ./bin/tparse
	go test -json ./... | tparse -all

acceptance-test: build
	docker-compose run tests

ci-test:
	go test -race -coverprofile=coverage.txt -covermode=atomic ./...

lint: ./bin/golangci-lint
	golangci-lint run

release: clean
	cd cmd/$(NAME) ; $(GORELEASER_BIN) $(PUBLISH_PARAM)

update:
	go get -u
	go mod tidy
	make test
	go mod tidy
