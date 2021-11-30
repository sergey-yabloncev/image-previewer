BIN := "./bin/app"
GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.37.0

lint: install-lint-deps
	golangci-lint run ./...

format:
	gofumpt -l -w .

build:
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd/app

test:
	go test -race -count 100 ./internal/...

integration-test-run:
	docker-compose -f docker-compose.yaml -f docker-compose-test.yaml up -d
	docker-compose -f docker-compose.yaml -f docker-compose-test.yaml run integration-tests
	docker-compose -f docker-compose.yaml -f docker-compose-test.yaml down

down:
	docker-compose down

run:
	docker-compose build
	docker-compose up -d

.PHONY: run test lint format down