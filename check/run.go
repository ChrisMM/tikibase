package check

import (
	"strings"

	"github.com/kevgo/tikibase/domain"
	"github.com/kevgo/tikibase/helpers"
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
	links := []domain.Link{}
	for d := range docs {
		docLinks := docs[d].Links()
		for l := range docLinks {
			// ignore external links
			if helpers.IsURL(docLinks[l].Target()) {
				continue
			}
			links = append(links, docLinks[l])
		}
	}

	// determine broken links
	for l := range links {
		target := links[l].Target()
		docFileName := links[l].SourceSection().Document().FileName()
		if isBrokenLink(target, docFileName, linkTargets) {
			brokenLinks = append(brokenLinks, BrokenLink{docFileName, target})
		}
	}

	// determine non-linked resources
	resourceFileNames := resourceFiles.FileNames()
	for f := range resourceFileNames {
		for l := range links {
			if links[l].Target() == resourceFileNames[f] {
				// TODO: this is wrong, make better feature and use LinkCollection.Contains here
				continue
			}
		}
		nonLinkedResources = append(nonLinkedResources, resourceFileNames[f])
	}

	return brokenLinks, duplicates, nonLinkedResources, err
}

func isBrokenLink(target string, filename domain.DocumentFilename, targets linkTargetCollection) bool {
	if strings.HasPrefix(target, "#") {
		return !targets.Contains(string(filename) + target)
	}
	return !targets.Contains(target)
}
