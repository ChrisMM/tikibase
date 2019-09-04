package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"github.com/kevgo/exit"
	"github.com/kevgo/tikibase/src/mentions"
)

//nolint:unused
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
	return ioutil.WriteFile(path.Join(w.root, filename), []byte(content.Content), 0644)
}

func (w *workspaceFeature) runMentions() error {
	return mentions.Run(w.root)
}

func (w *workspaceFeature) shouldContainFileWithContent(filename string, content *gherkin.DocString) error {
	data, err := ioutil.ReadFile(path.Join(w.root, filename))
	exit.IfWrapf(err, "Cannot find file '%s' in workspace", filename)
	text := string(data)
	if strings.Compare(text, content.Content) != 0 {
		return fmt.Errorf("text is different:\n\nEXPECTED:\n%s\n\nACTUAL:\n%s", content.Content, text)
	}
	return nil
}

//nolint:deadcode,unused
func FeatureContext(s *godog.Suite) {
	workspace := &workspaceFeature{}
	s.BeforeScenario(workspace.createWorkspace)
	s.Step(`^the workspace contains file "([^"]*)" with content:$`, workspace.containsFileWithContent)
	s.Step(`^running Mentions$`, workspace.runMentions)
	s.Step(`^the workspace should contain the file "([^"]*)" with content:$`, workspace.shouldContainFileWithContent)
}
