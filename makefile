.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

vet: fmt
	go vet ./...
.PHONY:vet

build: vet
	go build hello.go
.PHONY:build
