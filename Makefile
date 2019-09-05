build: clean    # builds for the current platform
	@go install

clean:   # Removes all build artifacts
	@go clean -i

cuke:  # runs the feature specs
	godog --concurrency=$(shell nproc --all) --format=progress

help:   # prints all make targets
	@cat Makefile | grep '^[^ ]*:' | grep -v '.PHONY' | grep -v help | sed 's/:.*#/#/' | column -s "#" -t

lint:  # runs all linters
	@golangci-lint run --enable-all -D lll

mentions:  # runs the 'mentions' command
	@go run main.go mentions

run:  # runs the command
	@go run main.go

test:  # runs the unit tests
	@go test ./...
