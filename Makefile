.DEFAULT_GOAL := dev

build:   # builds for the current platform
	@env CGO_ENABLED=0 go install
# Need to disable CGO here so that the binary is statically linked
# and will work on CI servers without having to install ld.so.

clean:   # Removes all build artifacts
	@go clean -i

cuke:  # runs the feature specs
	@env godog --format=progress

cuke-parallel:  # runs the feature specs
	@godog --concurrency=$(shell nproc --all) --format=progress

fix:   # fixes all auto-correctable issues
	@find . -name '*.go' | grep -v vendor | xargs gofmt -l -s -w
.PHONY: fix

help:   # prints all make targets
	@cat Makefile | grep '^[^ ]*:' | grep -v '.PHONY' | grep -v help | sed 's/:.*#/#/' | column -s "#" -t

lint:  # runs all linters
	@golangci-lint run --enable-all -D lll -D godox -D wsl -D whitespace

ps: fix test # "pitstop" command, run after making changes to the code

stats:  # shows code statistics
	@find . -type f | grep -v 'node_modules' | grep -v '\./.git/' | grep -v '\./vendor/' | xargs scc

test:  # runs all tests
	@go test ./... &
	@golangci-lint run --enable-all -D lll -D godox -D wsl -D whitespace &
	@godog --concurrency=$(shell nproc --all) --format=progress
.PHONY: test

unit:  # runs the unit tests
	@go test ./...

update:  # updates dependencies
	go get -u -t ./...

vendor:  # create/fix the vendor infrastructure
	go mod vendor
