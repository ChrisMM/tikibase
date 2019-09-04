package main

import (
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
)

func FeatureContext(s *godog.Suite) {
	s.Step(`^the workspace contains file "([^"]*)" with content:$`, func(arg1 string, arg2 *gherkin.DocString) error {
		return godog.ErrPending
	})
	s.Step(`^running Mentions$`, func() error {
		return godog.ErrPending
	})
	s.Step(`^the workspace should contain the file "([^"]*)" with content:$`, func(arg1 string, arg2 *gherkin.DocString) error {
		return godog.ErrPending
	})
}
