package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"

	"github.com/google/go-cmp/cmp"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/gherkin"
	"github.com/kevgo/tikibase/list"
	"github.com/kevgo/tikibase/occurrences"
	"github.com/kevgo/tikibase/test"
)

//nolint:unused
type workspaceFeature struct {

	// the root directory of this workspace
	root string

	// cache of the file contents
	fileContents map[string]string

	listResult []string
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
		return fmt.Errorf("no cached content for file %q found", filename)
	}
	data, err := ioutil.ReadFile(path.Join(w.root, filename))
	if err != nil {
		return fmt.Errorf("Cannot find file %q in workspace: %w", filename, err)
	}
	actual := string(data)
	if diff := cmp.Diff(expected, actual); diff != "" {
		return fmt.Errorf("mismatching content for file %s: \n%s", filename, diff)
	}
	return nil
}

func (w *workspaceFeature) itFinds(table *gherkin.DataTable) error {
	if len(table.Rows) != len(w.listResult) {
		return fmt.Errorf("expected %d results, got %d", len(table.Rows), len(w.listResult))
	}
	for i := range table.Rows {
		expected := table.Rows[i].Cells[0].Value
		actual := w.listResult[i]
		if actual != expected {
			return fmt.Errorf("mismatching entry %d: expected %s, got %s", i, expected, actual)
		}
	}
	return nil
}

func (w *workspaceFeature) listing(argument string) error {
	var err error
	w.listResult, err = list.Run(w.root, []string{argument})
	return err
}

func (w *workspaceFeature) runOccurrences() error {
	return occurrences.Run(w.root)
}

func (w *workspaceFeature) shouldContainFileWithContent(filename string, content *gherkin.DocString) error {
	data, err := ioutil.ReadFile(path.Join(w.root, filename))
	if err != nil {
		return fmt.Errorf("Cannot find file %q in workspace: %w", filename, err)
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
	s.Step(`^it finds:$`, workspace.itFinds)
	s.Step(`^listing "([^"]+)"$`, workspace.listing)
	s.Step(`^running Occurrences$`, workspace.runOccurrences)
	s.Step(`^the workspace contains a binary file "([^"]*)"$`, workspace.containsBinaryFile)
	s.Step(`^the workspace contains file "([^"]*)" with content:$`, workspace.containsFileWithContent)
	s.Step(`^the workspace should contain the file "([^"]*)" with content:$`, workspace.shouldContainFileWithContent)
}
