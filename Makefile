build: clean    # builds for the current platform
	@go install

clean:   # Removes all build artifacts
	@go clean -i

cuke:  # runs the feature specs
	@godog --format=progress

cuke-parallel:  # runs the feature specs
	godog --concurrency=$(shell nproc --all) --format=progress

fix:   # fixes all auto-correctable issues
	find . -name '*.go' | grep -v vendor | xargs gofmt -l -s -w

help:   # prints all make targets
	@cat Makefile | grep '^[^ ]*:' | grep -v '.PHONY' | grep -v help | sed 's/:.*#/#/' | column -s "#" -t

lint:  # runs all linters
	@golangci-lint run --enable-all -D lll -D godox -D wsl

occurrences:  # runs the 'occurrences' command
	@go run main.go occurrences

run:  # runs the command
	@go run main.go

test:  # runs all tests
	@go test ./... &
	@golangci-lint run --enable-all -D lll -D godox -D wsl &
	@godog --concurrency=$(shell nproc --all) --format=progress
.PHONY: test

unit:  # runs the unit tests
	@go test ./...
