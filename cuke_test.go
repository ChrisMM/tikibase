package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/google/go-cmp/cmp"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"github.com/kevgo/tikibase/mentions"
	"github.com/pkg/errors"
)

//nolint:unused
type workspaceFeature struct {
	root string
}

func (w *workspaceFeature) createWorkspace(arg interface{}) {
	_, err := os.Stat("tmp")
	if os.IsNotExist(err) {
		err = os.Mkdir("tmp", os.ModeDir|0777)
		if err != nil {
			log.Fatalf("cannot create root tmp directory: %s", err.Error())
		}
	}
	w.root, err = ioutil.TempDir("tmp", "")
	if err != nil {
		log.Fatalf("cannot create workspace: %s", err.Error())
	}
}

func (w *workspaceFeature) containsFileWithContent(filename string, content *gherkin.DocString) error {
	return ioutil.WriteFile(path.Join(w.root, filename), []byte(content.Content), 0644)
}

func (w *workspaceFeature) runMentions() error {
	return mentions.Run(w.root)
}

func (w *workspaceFeature) shouldContainFileWithContent(filename string, content *gherkin.DocString) error {
	data, err := ioutil.ReadFile(path.Join(w.root, filename))
	if err != nil {
		return errors.Wrapf(err, "Cannot find file '%s' in workspace", filename)
	}
	actual := string(data)
	expected := content.Content
	if strings.Compare(actual, expected) != 0 {
		diff := cmp.Diff(expected, actual)
		if diff != "" {
			return fmt.Errorf("mismatching content for file %s: \n%s", filename, diff)
		}
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
