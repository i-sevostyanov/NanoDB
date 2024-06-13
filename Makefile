.PHONY: run lint test coverage coverage-report generate build

run:
	@go run cmd/shell/main.go

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

build:
	@go build -trimpath -ldflags "-s -w" -o ./bin/shell cmd/shell/main.go
