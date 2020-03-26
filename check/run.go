package check

import (
	"fmt"

	"github.com/kevgo/tikibase/domain"
)

// Result contains the outcome of a TikiBase check.
type Result struct {
	BrokenLinks                []brokenLink
	Duplicates                 []string
	DocumentsWithEmptySections []string
	MixedCapSections           [][]string
	NonLinkedResources         []string
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
	internalLinks, _ := docs.Links()

	result.BrokenLinks = findBrokenLinks(docs, resourceFiles, internalLinks)
	result.Duplicates, err = findDuplicateTargets(docs)
	if err != nil {
		return result, fmt.Errorf("cannot determine duplicate link targets: %w", err)
	}
	result.NonLinkedResources = findNonLinkedResources(resourceFiles, internalLinks)
	result.DocumentsWithEmptySections, err = docsWithEmptySections(docs)
	if err != nil {
		return result, fmt.Errorf("cannot determine documents with empty sections: %w", err)
	}
	result.MixedCapSections, err = findMixedCapSections(docs)
	if err != nil {
		return result, fmt.Errorf("cannot find mixed-cap sections: %w", err)
	}
	return result, nil
}
