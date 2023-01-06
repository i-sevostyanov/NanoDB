.PHONY: lint test coverage generate

lint:
	@golangci-lint run

test:
	@go test ./... -race

coverage:
	@go test ./... -race -coverprofile=coverage.txt -covermode=atomic

coverage-report:
	@go tool cover -html=coverage.txt

generate:
	@go generate ./...
