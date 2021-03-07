.PHONY: lint test coverage

lint:
	@golangci-lint run

test:
	@go test ./...

coverage:
	@go test ./... -race -coverprofile=coverage.txt -covermode=atomic
