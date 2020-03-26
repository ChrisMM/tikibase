package check

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kevgo/tikibase/domain"
)

// linkTargets contains all possible link targets within this Tikibase.
//
// Examples: `1.md`, `1.md#foo`
type linkTargets map[string]struct{}

func findLinkTargets(docs domain.Documents, resources domain.ResourceFiles) (result linkTargets, err error) {
	result = make(linkTargets)

	// add links targets for documents
	for i := range docs {

		// add target for the document itself
		result.Add(docs[i].FileName())

		// add target for the sections in the document
		sections := docs[i].AllSections()
		for s := range sections {
			linkTarget, err := sections[s].LinkTarget()
			if err != nil {
				return result, fmt.Errorf("cannot determine link targets in document %q: %w", docs[i].FileName(), err)
			}
			result.Add(linkTarget)
		}
	}

	// add link targets for resources
	for r := range resources {
		result.Add(resources[r])
	}
	return result, nil
}

func (ltc linkTargets) Add(linkTarget string) {
	ltc[linkTarget] = struct{}{}
}

func (ltc linkTargets) Contains(linkTarget string) bool {
	_, ok := ltc[linkTarget]
	return ok
}

func (ltc linkTargets) String() (result string) {
	values := []string{}
	for value := range ltc {
		values = append(values, value)
	}
	sort.Strings(values)
	return "[" + strings.Join(values, ", ") + "]"
}

// ScaffoldLinkTargets provides linkTargets instances for testing.
func scaffoldLinkTargets(targets []string) linkTargets {
	result := make(linkTargets)
	for t := range targets {
		result.Add(targets[t])
	}
	return result
}
