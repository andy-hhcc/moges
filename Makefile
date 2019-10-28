PROJECT := moges
PKGS := $(shell go list ./... | grep -v /vendor)
EXECUTABLE := moges
GIT_COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null)

.PHONY: install build fmt

GO_LDFLAGS := -X $(shell go list ./$(PACKAGE)).GitCommit=$(GIT_COMMIT)

all: build

install:
	go mod vendor

clean:
	rm -rf bin

build:
	@go mod vendor
	@mkdir -p bin/conf && cp -vr conf bin/
	@go build -i -ldflags "$(GO_LDFLAGS)" -o bin/$(EXECUTABLE)

fmt:
		go fmt ./...
