.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

test:
	go test ./...
.PHONY:test

vtest:
	go test ./... -v
.PHONY:vtest
