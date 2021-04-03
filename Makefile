.PHONY: lint test coverage generate

lint:
	@golangci-lint run

test:
	@go test ./...

coverage:
	@go test ./... -race -coverprofile=coverage.txt -covermode=atomic

generate:
	@go generate ./...
