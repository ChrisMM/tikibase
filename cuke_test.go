package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"

	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v10"
	"github.com/google/go-cmp/cmp"
	"github.com/kevgo/tikibase/check"
	"github.com/kevgo/tikibase/find"
	"github.com/kevgo/tikibase/fix"
	"github.com/kevgo/tikibase/linkify"
	"github.com/kevgo/tikibase/stats"
	"github.com/kevgo/tikibase/test"
)

//nolint:unused
type workspaceFeature struct {

	// the root directory of this workspace
	root string

	// cache of the file contents
	fileContents map[string]string

	findResult []string

	checkResult check.Result

	statisticsResult stats.Result
}

func (w *workspaceFeature) checkingTheTikiBase() (err error) {
	w.checkResult, err = check.Run(w.root)
	return err
}

func (w *workspaceFeature) containsBinaryFile(filename string) error {
	return test.CreateBinaryFile(filepath.Join(w.root, filename))
}

func (w *workspaceFeature) containsFileWithContent(filename string, content *messages.PickleStepArgument_PickleDocString) error {
	w.fileContents[filename] = content.Content + "\n"
	return ioutil.WriteFile(filepath.Join(w.root, filename), []byte(content.Content+"\n"), 0644)
}

func (w *workspaceFeature) createWorkspace(arg *messages.Pickle) {
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
	data, err := ioutil.ReadFile(filepath.Join(w.root, filename))
	if err != nil {
		return fmt.Errorf("Cannot find file %q in workspace: %w", filename, err)
	}
	actual := string(data)
	if diff := cmp.Diff(expected, actual); diff != "" {
		return fmt.Errorf("mismatching content for file %s: \n%s", filename, diff)
	}
	return nil
}

func (w *workspaceFeature) finding(argument string) error {
	var err error
	w.findResult, err = find.Run(w.root, []string{argument})
	return err
}

func (w *workspaceFeature) fixingTheTikiBase() error {
	return fix.Run(w.root)
}

func (w *workspaceFeature) itFinds(table *messages.PickleStepArgument_PickleTable) error {
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

func (w *workspaceFeature) itFindsNoErrors() error {
	if len(w.checkResult.BrokenLinks) > 0 {
		msg := fmt.Sprintf("Found %d broken links: \n", len(w.checkResult.BrokenLinks))
		for i := range w.checkResult.BrokenLinks {
			msg += fmt.Sprintf("- %s: %q", w.checkResult.BrokenLinks[i].Filename, w.checkResult.BrokenLinks[i].Link)
		}
		return fmt.Errorf(msg)
	}
	if len(w.checkResult.NonLinkedResources) > 0 {
		msg := fmt.Sprintf("Found %d non-linked resources: \n", len(w.checkResult.NonLinkedResources))
		for i := range w.checkResult.NonLinkedResources {
			msg += fmt.Sprintf("- %s", w.checkResult.NonLinkedResources[i])
		}
		return fmt.Errorf(msg)
	}
	if len(w.checkResult.Duplicates) > 0 {
		msg := fmt.Sprintf("Found %d duplicates: \n", len(w.checkResult.Duplicates))
		for i := range w.checkResult.Duplicates {
			msg += fmt.Sprintf("- %s", w.checkResult.Duplicates[i])
		}
		return fmt.Errorf(msg)
	}
	if len(w.checkResult.MixedCapSections) > 0 {
		msg := fmt.Sprintf("Found %d mixed cap sections: \n", len(w.checkResult.MixedCapSections))
		for i := range w.checkResult.MixedCapSections {
			msg += fmt.Sprintf("- %v", w.checkResult.MixedCapSections[i])
		}
		return fmt.Errorf(msg)
	}
	return nil
}

func (w *workspaceFeature) itFindsTheBrokenLinks(expected *messages.PickleStepArgument_PickleTable) error {
	if len(w.checkResult.BrokenLinks) != len(expected.Rows)-1 {
		return fmt.Errorf("expected %d broken links but got %d", len(expected.Rows)-1, len(w.checkResult.BrokenLinks))
	}
	for i := 0; i < len(w.checkResult.BrokenLinks); i++ {
		expectedFile := expected.Rows[i+1].Cells[0].Value
		expectedLink := expected.Rows[i+1].Cells[1].Value
		actualFile := w.checkResult.BrokenLinks[i].Filename
		actualLink := w.checkResult.BrokenLinks[i].Link
		if expectedFile != actualFile || expectedLink != actualLink {
			return fmt.Errorf("expected file %q to contain broken link %q, instead got file %q with broken link %q", expectedFile, expectedLink, actualFile, actualLink)
		}
	}
	return nil
}

func (w *workspaceFeature) itFindsTheDuplicates(table *messages.PickleStepArgument_PickleTable) error {
	expected := []string{}
	for i := range table.Rows {
		expected = append(expected, table.Rows[i].Cells[0].Value)
	}
	if !reflect.DeepEqual(w.checkResult.Duplicates, expected) {
		return fmt.Errorf("expected %v, got %v", expected, w.checkResult.Duplicates)
	}
	return nil
}

func (w *workspaceFeature) itFindsTheNonlinkedResources(table *messages.PickleStepArgument_PickleTable) error {
	expected := make([]string, len(table.Rows))
	for i := range table.Rows {
		expected[i] = table.Rows[i].Cells[0].Value
	}
	if !reflect.DeepEqual(expected, w.checkResult.NonLinkedResources) {
		return fmt.Errorf("expected %s, got %s", expected, w.checkResult.NonLinkedResources)
	}
	return nil
}

func (w *workspaceFeature) itFindsTheSectionTypes(table *messages.PickleStepArgument_PickleTable) error {
	expected := make([]string, len(table.Rows))
	for i := range table.Rows {
		expected[i] = table.Rows[i].Cells[0].Value
	}
	if !reflect.DeepEqual(expected, w.statisticsResult.SectionTypes) {
		return fmt.Errorf("expected %s, got %s", expected, w.statisticsResult.SectionTypes)
	}
	return nil
}

func (w *workspaceFeature) itFindsTheseSectionsWithMixedCapitalization(table *messages.PickleStepArgument_PickleTable) error {
	for r := range table.Rows {
		expected := strings.Split(table.Rows[r].Cells[0].Value, ", ")
		if !reflect.DeepEqual(w.checkResult.MixedCapSections[r], expected) {
			return fmt.Errorf("expected [%s], found [%s]", strings.Join(expected, ","), strings.Join(w.checkResult.MixedCapSections[r], ","))
		}
	}
	return nil
}

func (w *workspaceFeature) itProvidesTheStatistics(table *messages.PickleStepArgument_PickleTable) error {
	// check docs count
	expectedDocsCount, err := strconv.Atoi(table.Rows[0].Cells[1].Value)
	if err != nil {
		return fmt.Errorf("expected docs count is not a number: %w", err)
	}
	if w.statisticsResult.DocsCount != expectedDocsCount {
		return fmt.Errorf("expected %d docs, got %d", expectedDocsCount, w.statisticsResult.DocsCount)
	}
	// check sections count
	expectedSectionsCount, err := strconv.Atoi(table.Rows[1].Cells[1].Value)
	if err != nil {
		return fmt.Errorf("expected sections count is not a number: %w", err)
	}
	if w.statisticsResult.SectionsCount != expectedSectionsCount {
		return fmt.Errorf("expected %d sections, got %d", expectedSectionsCount, w.statisticsResult.SectionsCount)
	}
	// check links count
	expectedLinksCount, err := strconv.Atoi(table.Rows[2].Cells[1].Value)
	if err != nil {
		return fmt.Errorf("expected links count is not a number: %w", err)
	}
	if w.statisticsResult.LinksCount != expectedLinksCount {
		return fmt.Errorf("expected %d links, got %d", expectedLinksCount, w.statisticsResult.LinksCount)
	}
	// check resources count
	expectedResourcesCount, err := strconv.Atoi(table.Rows[3].Cells[1].Value)
	if err != nil {
		return fmt.Errorf("expected resources count is not a number: %w", err)
	}
	if w.statisticsResult.ResourcesCount != expectedResourcesCount {
		return fmt.Errorf("expected %d links, got %d", expectedResourcesCount, w.statisticsResult.ResourcesCount)
	}
	return nil
}

func (w *workspaceFeature) linkify() (err error) {
	return linkify.Run(w.root, false)
}

func (w *workspaceFeature) runningStatistics() error {
	var err error
	w.statisticsResult, err = stats.Run(w.root)
	return err
}

func (w *workspaceFeature) shouldContainFileWithContent(filename string, text *messages.PickleStepArgument_PickleDocString) error {
	data, err := ioutil.ReadFile(filepath.Join(w.root, filename))
	if err != nil {
		return fmt.Errorf("Cannot find file %q in workspace: %w", filename, err)
	}
	actual := string(data)
	expected := text.Content + "\n"
	if actual != expected {
		return fmt.Errorf("mismatching content for file %s:\n- have: %s\n- want: %s", filename, actual, expected)
	}
	return nil
}

//nolint:deadcode,unused
func FeatureContext(s *godog.Suite) {
	wf := &workspaceFeature{fileContents: make(map[string]string)}
	s.BeforeScenario(wf.createWorkspace)
	s.Step(`^checking the TikiBase$`, wf.checkingTheTikiBase)
	s.Step(`^file "([^"]*)" is unchanged$`, wf.fileIsUnchanged)
	s.Step(`^linkifying$`, wf.linkify)
	s.Step(`^it finds:$`, wf.itFinds)
	s.Step(`^it finds no errors$`, wf.itFindsNoErrors)
	s.Step(`^it finds the broken links:$`, wf.itFindsTheBrokenLinks)
	s.Step(`^it finds the duplicates:$`, wf.itFindsTheDuplicates)
	s.Step(`^it finds the non-linked resources:$`, wf.itFindsTheNonlinkedResources)
	s.Step(`^it finds the section types:$`, wf.itFindsTheSectionTypes)
	s.Step(`^it finds these sections with mixed capitalization:$`, wf.itFindsTheseSectionsWithMixedCapitalization)
	s.Step(`^it provides the statistics:$`, wf.itProvidesTheStatistics)
	s.Step(`^finding "([^"]+)"$`, wf.finding)
	s.Step(`^fixing the TikiBase$`, wf.fixingTheTikiBase)
	s.Step(`^running Statistics$`, wf.runningStatistics)
	s.Step(`^the workspace contains a binary file "([^"]*)"$`, wf.containsBinaryFile)
	s.Step(`^the workspace contains file "([^"]*)" with content:$`, wf.containsFileWithContent)
	s.Step(`^the workspace should contain the file "([^"]*)" with content:$`, wf.shouldContainFileWithContent)
}
