SHELL := /bin/bash
export GOBIN = $(shell pwd)/bin
VERSION := 0.1 # We can use git tag for a more consistent results $(shell git tag -l $(VERSION))
BINARY := gramaddict_helper
BUILD_PATH ?= $(shell go env GOPATH)/bin
OS := $(shell uname -s | tr '[:upper:]' '[:lower:]')

define HELP


$(BINARY) v$(VERSION) Makefile
=================================

## Build target

- binary:                It will build $(BINARY) for the current system (by default in $(GOBIN)).

## Development targets

- lint:                   Runs golint and gofmt against all of the packages.
- format:                 Formats the codebase according to gofmt and goimports.
- unit:                   Runs unit tests.
- docker:                 Build docker image with the compiled binary.

## Release targets

- release:                Releases the package to docker repository.

endef
export HELP

.DEFAULT: help
.PHONY: help
help:
	@ echo "$$HELP"

include build/Makefile.deps

.PHONY: binary
binary:
	@ echo "Building linux binary into $(GOBIN)/$(BINARY)"
	@ echo "Binary will be based on $(OS) $(ARCH)"
	@ BINARY=$(BINARY) GOOS=$(OS) GOARCH=$(ARCH) go build -o $(GOBIN)/$(BINARY)

.PHONY: release
release: docker
	@ echo "------------------------------------"
	@ echo "Pushing docker image to repo..."

.PHONY: docker
docker: binary
	@ echo "Dockerizing..."
	@ docker build -t yaniv:$(VERSION) .

.PHONY: unit
unit:
	@ echo "Running unit tests for $(BINARY)..."
	go test

.PHONY: lint
lint: deps
	@ echo " Running linters..."
	@ $(GOBIN)/golint -set_exit_status ./...
	@ $(GOBIN)/golangci-lint run --timeout 2m
	@ echo "Done."

.PHONY: format
format: deps
	@ echo " Formatting/auto-fixing Go files..."
	@ $(GOBIN)/golangci-lint run --fix
	@ echo "Done."

.PHONY: clean
clean:
	@ rm -rf bin