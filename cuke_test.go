package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"reflect"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/gherkin"
	"github.com/google/go-cmp/cmp"
	"github.com/kevgo/tikibase/check"
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

	brokenLinks []check.BrokenLink

	duplicates []string
}

func (w *workspaceFeature) checkingTheLinks() (err error) {
	w.brokenLinks, w.duplicates, err = check.Run(w.root)
	return err
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

func (w *workspaceFeature) itFindsTheDuplicates(table *gherkin.DataTable) error {
	expected := []string{}
	for i := range table.Rows {
		expected = append(expected, table.Rows[i].Cells[0].Value)
	}
	if !reflect.DeepEqual(w.duplicates, expected) {
		return fmt.Errorf("expected %v, got %v", expected, w.duplicates)
	}
	return nil
}

func (w *workspaceFeature) itFindsNoErrors() error {
	if len(w.brokenLinks) == 0 {
		return nil
	}
	msg := fmt.Sprintf("Found %d errors: \n", len(w.brokenLinks))
	for i := range w.brokenLinks {
		msg += fmt.Sprintf("- file %q contains broken link %q", w.brokenLinks[i].Filename, w.brokenLinks[i].Link)
	}
	return fmt.Errorf(msg)
}

func (w *workspaceFeature) itFindsTheBrokenLinks(expected *gherkin.DataTable) error {
	if len(w.brokenLinks) != len(expected.Rows)-1 {
		return fmt.Errorf("expected %d broken links but got %d", len(expected.Rows)-1, len(w.brokenLinks))
	}
	for i := 0; i < len(w.brokenLinks); i++ {
		expectedFile := expected.Rows[i+1].Cells[0].Value
		expectedLink := expected.Rows[i+1].Cells[1].Value
		actualFile := string(w.brokenLinks[i].Filename)
		actualLink := w.brokenLinks[i].Link
		if expectedFile != actualFile || expectedLink != actualLink {
			return fmt.Errorf("expected file %q to contain broken link %q, instead got file %q with broken link %q", expectedFile, expectedLink, actualFile, actualLink)
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
	_, _, _, _, err := occurrences.Run(w.root)
	return err
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
	s.Step(`^checking the links$`, workspace.checkingTheLinks)
	s.Step(`^file "([^"]*)" is unchanged$`, workspace.fileIsUnchanged)
	s.Step(`^it finds:$`, workspace.itFinds)
	s.Step(`^it finds no errors$`, workspace.itFindsNoErrors)
	s.Step(`^it finds the broken links:$`, workspace.itFindsTheBrokenLinks)
	s.Step(`^it finds the duplicates:$`, workspace.itFindsTheDuplicates)
	s.Step(`^listing "([^"]+)"$`, workspace.listing)
	s.Step(`^running Occurrences$`, workspace.runOccurrences)
	s.Step(`^the workspace contains a binary file "([^"]*)"$`, workspace.containsBinaryFile)
	s.Step(`^the workspace contains file "([^"]*)" with content:$`, workspace.containsFileWithContent)
	s.Step(`^the workspace should contain the file "([^"]*)" with content:$`, workspace.shouldContainFileWithContent)
}
