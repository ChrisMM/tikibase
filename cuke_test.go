package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"

	"github.com/google/go-cmp/cmp"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"github.com/kevgo/tikibase/mentions"
	"github.com/kevgo/tikibase/test"
	"github.com/pkg/errors"
)

//nolint:unused
type workspaceFeature struct {

	// the root directory of this workspace
	root string

	// cache of the file contents
	fileContents map[string]string
}

func (w *workspaceFeature) containsBinaryFile(filename string) error {
	return test.CreateBinaryFile(path.Join(w.root, filename))
}

func (w *workspaceFeature) containsFileWithContent(filename string, content *gherkin.DocString) error {
	w.fileContents[filename] = content.Content + "\n"
	return ioutil.WriteFile(path.Join(w.root, filename), []byte(content.Content+"\n"), 0644)
}

func (w *workspaceFeature) createWorkspace(arg interface{}) {
	var err error
	w.root, err = ioutil.TempDir("", "")
	if err != nil {
		log.Fatalf("cannot create workspace: %s", err.Error())
	}
}

func (w *workspaceFeature) fileIsUnchanged(filename string) error {
	expected, exists := w.fileContents[filename]
	if !exists {
		return fmt.Errorf("no cached content for file '%s' found", filename)
	}
	data, err := ioutil.ReadFile(path.Join(w.root, filename))
	if err != nil {
		return errors.Wrapf(err, "Cannot find file '%s' in workspace", filename)
	}
	actual := string(data)
	if diff := cmp.Diff(expected, actual); diff != "" {
		return fmt.Errorf("mismatching content for file %s: \n%s", filename, diff)
	}
	return nil
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
	expected := content.Content + "\n"
	if diff := cmp.Diff(expected, actual); diff != "" {
		return fmt.Errorf("mismatching content for file %s: \n%s", filename, diff)
	}
	return nil
}

//nolint:deadcode,unused
func FeatureContext(s *godog.Suite) {
	workspace := &workspaceFeature{fileContents: make(map[string]string)}
	s.BeforeScenario(workspace.createWorkspace)
	s.Step(`^file "([^"]*)" is unchanged$`, workspace.fileIsUnchanged)
	s.Step(`^running Mentions$`, workspace.runMentions)
	s.Step(`^the workspace contains a binary file "([^"]*)"$`, workspace.containsBinaryFile)
	s.Step(`^the workspace contains file "([^"]*)" with content:$`, workspace.containsFileWithContent)
	s.Step(`^the workspace should contain the file "([^"]*)" with content:$`, workspace.shouldContainFileWithContent)
}
