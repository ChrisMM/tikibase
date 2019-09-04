build: clean    # builds for the current platform
	@go install

clean:   # Removes all build artifacts
	@go clean -i

cuke:  # runs the feature specs
	@godog

help:   # prints all make targets
	@cat Makefile | grep '^[^ ]*:' | grep -v '.PHONY' | grep -v help | sed 's/:.*#/#/' | column -s "#" -t
