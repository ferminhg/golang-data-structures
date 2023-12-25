.DEFAULT_GOAL := build
.PHONY:fmt lint vet test vtest

fmt:
	go fmt ./...

lint: fmt
	golangci-lint run

vet: fmt
	go vet ./...

test:
	go test ./...

vtest:
	go test ./... -v
