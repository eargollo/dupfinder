VERSION=$(shell git describe --tags)

.PHONY: build
build:
	CGO_ENABLED=0 go build -ldflags "-extldflags=-static" .

.PHONY: test
test:
	go test -cover ./...

.PHONY: lint
lint: lint-code lint-security lint-vulnerability

.PHONY: lint-code 
lint-code:
	golangci-lint run

.PHONY: lint-security
lint-security:
	gosec ./...
	
.PHONY: lint-vulnerability
lint-vulnerability:
	govulncheck ./...

.PHONY: outdated
outdated:
	@go list -u -m -f '{{if not .Indirect}}{{.}}{{end}}' all | grep -F '[' || true

.PHONY: cover
cover:
	go test -coverprofile=coverage.out -covermode=count  ./...
	@go tool cover -func=coverage.out | grep total | grep -Eo '[0-9]+\.[0-9]+'
