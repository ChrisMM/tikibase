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
	resourceFileNames := resourceFiles.FileNames()
	for f := range resourceFileNames {
		if !internalLinks.HasLinkTo(resourceFileNames[f]) {
			nonLinkedResources = append(nonLinkedResources, resourceFileNames[f])
		}
	}

	return brokenLinks, duplicates, nonLinkedResources, err
}

func isBrokenLink(target string, filename domain.DocumentFilename, targets linkTargetCollection) bool {
	if strings.HasPrefix(target, "#") {
		return !targets.Contains(string(filename) + target)
	}
	return !targets.Contains(target)
}
