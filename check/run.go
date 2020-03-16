package check

import (
	"fmt"
	"strings"

	"github.com/kevgo/tikibase/domain"
	"github.com/kevgo/tikibase/helpers"
)

// Result contains the outcome of a TikiBase check.
type Result struct {
	BrokenLinks        []BrokenLink
	Duplicates         []string
	MixedCapSections   [][]string
	NonLinkedResources []string
}

// Run executes the "check" command.
func Run(dir string) (result Result, err error) {
	tikibase, err := domain.NewTikiBase(dir)
	if err != nil {
		return
	}
	docFiles, resourceFiles, err := tikibase.Files()
	if err != nil {
		return
	}
	docs, err := docFiles.Documents()
	if err != nil {
		return
	}
	var targets linkTargets
	targets, result.Duplicates, err = findLinkTargets(docs, resourceFiles)
	if err != nil {
		return
	}
	internalLinks, _ := docs.Links()

	// determine broken links
	for l := range internalLinks {
		target := internalLinks[l].Target()
		docFileName := internalLinks[l].SourceSection().Document().FileName()
		if isBrokenLink(target, docFileName, targets) {
			result.BrokenLinks = append(result.BrokenLinks, BrokenLink{docFileName, target})
		}
	}

	// determine non-linked resources
	for _, fileName := range resourceFiles.FileNames() {
		if !internalLinks.HasLinkTo(fileName) {
			result.NonLinkedResources = append(result.NonLinkedResources, fileName)
		}
	}

	// determine mixed cap sections
	titles := []string{}
	sections := docs.ContentSections()
	for s := range sections {
		title, err := sections[s].Title()
		if err != nil {
			return result, fmt.Errorf("cannot determine mixed cap sections: %w", err)
		}
		titles = append(titles, title)
	}
	titles = helpers.UniqueStrings(titles)
	helpers.SortCaseInsensitive(titles)
	result.MixedCapSections = FindGroups(titles)
	return result, err
}

func isBrokenLink(target string, filename string, targets linkTargets) bool {
	if strings.HasPrefix(target, "#") {
		return !targets.Contains(filename + target)
	}
	return !targets.Contains(target)
}
