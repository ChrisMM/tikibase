build: clean    # builds for the current platform
	@node_modules/.bin/tsc -p .

clean:   # Removes all build artifacts
	@rm -rf dist

cuke: build   # runs the feature specs
	@node_modules/.bin/cucumber-js --format progress

help:   # prints all make targets
	@cat Makefile | grep '^[^ ]*:' | grep -v '.PHONY' | grep -v help | sed 's/:.*#/#/' | column -s "#" -t
