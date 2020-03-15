package check

import (
	"strings"

	"github.com/kevgo/tikibase/domain"
)

// Run executes the "check" command.
func Run(dir string) (brokenLinks []BrokenLink, duplicates []string, nonLinkedResources []string, err error) {
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
	linkTargets, duplicates, err := findLinkTargets(docs, resourceFiles)
	if err != nil {
		return
	}

	// determine all links
	internalLinks, _ := docs.Links()

	// determine broken links
	for l := range internalLinks {
		target := internalLinks[l].Target()
		docFileName := internalLinks[l].SourceSection().Document().FileName()
		if isBrokenLink(target, docFileName, linkTargets) {
			brokenLinks = append(brokenLinks, BrokenLink{docFileName, target})
		}
	}

	// determine non-linked resources
	for _, fileName := range resourceFiles.FileNames() {
		if !internalLinks.HasLinkTo(fileName) {
			nonLinkedResources = append(nonLinkedResources, fileName)
		}
	}

	return brokenLinks, duplicates, nonLinkedResources, err
}

func isBrokenLink(target string, filename string, targets linkTargetCollection) bool {
	if strings.HasPrefix(target, "#") {
		return !targets.Contains(filename + target)
	}
	return !targets.Contains(target)
}
