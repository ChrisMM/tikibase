package check

import (
	"strings"

	"github.com/kevgo/tikibase/domain"
	"github.com/kevgo/tikibase/helpers"
)

// Run executes the "check" command.
func Run(dir string) (brokenLinks []BrokenLink, duplicates []string, err error) {
	tikibase, err := domain.NewTikiBase(dir)
	if err != nil {
		return
	}
	docs, err := tikibase.Documents()
	if err != nil {
		return
	}
	fileNames, err := tikibase.FileNames()
	if err != nil {
		return
	}
	linkTargets, duplicates, err := findLinkTargets(fileNames, docs)
	if err != nil {
		return
	}
	for i := range docs {
		brokenLinks = append(brokenLinks, checkDocLinks(docs[i], linkTargets)...)
	}
	return brokenLinks, duplicates, err
}

func checkDocLinks(doc *domain.Document, linkTargets linkTargetCollection) (brokenLinks []BrokenLink) {
	links := doc.Links()
	for i := range links {
		target := links[i].Target()
		// ignore external links
		if helpers.IsURL(target) {
			continue
		}
		if isBrokenLink(target, doc.FileName(), linkTargets) {
			brokenLinks = append(brokenLinks, BrokenLink{doc.FileName(), target})
		}
	}
	return brokenLinks
}

func isBrokenLink(target string, filename domain.DocumentFilename, targets linkTargetCollection) bool {
	if strings.HasPrefix(target, "#") {
		return !targets.Contains(string(filename) + target)
	}
	return !targets.Contains(target)
}
