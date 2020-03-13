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
	"github.com/kevgo/tikibase/find"
	"github.com/kevgo/tikibase/fix"
	"github.com/kevgo/tikibase/test"
)

//nolint:unused
type workspaceFeature struct {

	// the root directory of this workspace
	root string

	// cache of the file contents
	fileContents map[string]string

	findResult []string

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
	if len(table.Rows) != len(w.findResult) {
		return fmt.Errorf("expected %d results, got %d", len(table.Rows), len(w.findResult))
	}
	for i := range table.Rows {
		expected := table.Rows[i].Cells[0].Value
		actual := w.findResult[i]
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

func (w *workspaceFeature) finding(argument string) error {
	var err error
	w.findResult, err = find.Run(w.root, []string{argument})
	return err
}

func (w *workspaceFeature) runFix() error {
	_, _, _, _, err := fix.Run(w.root)
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
	wf := &workspaceFeature{fileContents: make(map[string]string)}
	s.BeforeScenario(wf.createWorkspace)
	s.Step(`^checking the links$`, wf.checkingTheLinks)
	s.Step(`^file "([^"]*)" is unchanged$`, wf.fileIsUnchanged)
	s.Step(`^it finds:$`, wf.itFinds)
	s.Step(`^it finds no errors$`, wf.itFindsNoErrors)
	s.Step(`^it finds the broken links:$`, wf.itFindsTheBrokenLinks)
	s.Step(`^it finds the duplicates:$`, wf.itFindsTheDuplicates)
	s.Step(`^finding "([^"]+)"$`, wf.finding)
	s.Step(`^running Fix$`, wf.runFix)
	s.Step(`^the workspace contains a binary file "([^"]*)"$`, wf.containsBinaryFile)
	s.Step(`^the workspace contains file "([^"]*)" with content:$`, wf.containsFileWithContent)
	s.Step(`^the workspace should contain the file "([^"]*)" with content:$`, wf.shouldContainFileWithContent)
}
