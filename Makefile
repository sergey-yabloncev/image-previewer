install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.37.0

lint:
	golangci-lint run ./...

format:
	gofumpt -l -w .

test:
	go test -race ./internal/...

down:
	docker-compose down

run:
	docker-compose build
	docker-compose up -d

.PHONY: run test lint format down