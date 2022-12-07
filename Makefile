PROJECT_NAME=go-backend
MODULE_NAME=go-backend

.DEFAULT_GOAL := build

.PHONY: build
build:
	@go build .

.PHONY: fmt
fmt:
	@go fmt ./...

.PHONY: test
test:
	@go test -v -coverprofile coverage.out ./...

.PHONY: coverage
coverage:
	@go tool cover -html=coverage.out

.PHONY: get
get:
	@go mod download

.PHONY: docker
docker:
	@docker build -t go-backend:latest .

