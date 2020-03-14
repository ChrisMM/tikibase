build:   # builds for the current platform
	@go install

clean:   # Removes all build artifacts
	@go clean -i

cuke:  # runs the feature specs
	@env GOFLAGS=-mod=vendor godog --format=progress

cuke-parallel:  # runs the feature specs
	@godog --concurrency=$(shell nproc --all) --format=progress

fix:   # fixes all auto-correctable issues
	find . -name '*.go' | grep -v vendor | xargs gofmt -l -s -w

help:   # prints all make targets
	@cat Makefile | grep '^[^ ]*:' | grep -v '.PHONY' | grep -v help | sed 's/:.*#/#/' | column -s "#" -t

lint:  # runs all linters
	@golangci-lint run --enable-all -D lll -D godox -D wsl -D whitespace

stats:  # shows code statistics
	@find . -type f | grep -v 'node_modules' | grep -v '\./.git/' | grep -v '\./vendor/' | xargs scc

test:  # runs all tests
	@go test ./... &
	@golangci-lint run --enable-all -D lll -D godox -D wsl -D whitespace &
	@godog --concurrency=$(shell nproc --all) --format=progress
.PHONY: test

unit:  # runs the unit tests
	@env GOFLAGS=-mod=vendor go test ./...
