package main

import (
	"io/ioutil"
	"os"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"github.com/kevgo/exit"
)

type workspaceFeature struct {
	root string
}

func (w *workspaceFeature) createWorkspace(arg interface{}) {
	_, err := os.Stat("tmp")
	if os.IsNotExist(err) {
		err = os.Mkdir("tmp", os.ModeDir|0777)
		exit.IfWrap(err, "cannot create root tmp directory")
	}
	w.root, err = ioutil.TempDir("tmp", "")
	exit.IfWrap(err, "cannot create workspace")
}

func (w *workspaceFeature) containsFileWithContent(filename string, content *gherkin.DocString) error {
	return godog.ErrPending
}

func (w *workspaceFeature) runningMentions() error {
	return godog.ErrPending
}

func (w *workspaceFeature) shouldContainFileWithContent(filename string, content *gherkin.DocString) error {
	return godog.ErrPending
}

//nolint:deadcode,unused
func FeatureContext(s *godog.Suite) {
	workspace := &workspaceFeature{}
	s.BeforeScenario(workspace.createWorkspace)
	s.Step(`^the workspace contains file "([^"]*)" with content:$`, workspace.containsFileWithContent)
	s.Step(`^running Mentions$`, workspace.runningMentions)
	s.Step(`^the workspace should contain the file "([^"]*)" with content:$`, workspace.shouldContainFileWithContent)
}
