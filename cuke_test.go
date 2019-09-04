package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"github.com/kevgo/tikibase/src/mentions"
	"github.com/pkg/errors"
	"github.com/sergi/go-diff/diffmatchpatch"
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
	text := string(data)
	if strings.Compare(text, content.Content) != 0 {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(text, content.Content, false)
		fmt.Println(dmp.DiffPrettyText(diffs))
		return fmt.Errorf("mismatching content for file %s", filename)
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
