package check

import (
	"strings"

	"github.com/kevgo/tikibase/domain"
)

// brokenLink describes a broken link.
type brokenLink struct {

	// path of the file containing the broken link
	Filename string

	// the Link that is broken
	Link string
}

func findBrokenLinks(docs domain.Documents, resourceFiles domain.ResourceFiles, internalLinks domain.Links) (result []brokenLink) {
	targets, err := findLinkTargets(docs, resourceFiles)
	if err != nil {
		return
	}
	for l := range internalLinks {
		target := internalLinks[l].Target()
		docFileName := internalLinks[l].SourceSection().Document().FileName()
		if isBrokenLink(target, docFileName, targets) {
			result = append(result, brokenLink{docFileName, target})
		}
	}
	return result
}

func isBrokenLink(target string, filename string, targets linkTargets) bool {
	if strings.HasPrefix(target, "#") {
		return !targets.Contains(filename + target)
	}
	return !targets.Contains(target)
}
