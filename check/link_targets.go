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

func findLinkTargets(docs domain.Documents, resources domain.ResourceFiles) (result linkTargets, duplicates []string, err error) {
	result = make(linkTargets)

	// add links targets for documents
	for i := range docs {

		// add target for the document itself
		err := result.Add(docs[i].FileName())
		if err != nil {
			duplicates = append(duplicates, docs[i].FileName())
		}

		// add target for the sections in the document
		sections := docs[i].AllSections()
		for s := range sections {
			linkTarget, err := sections[s].LinkTarget()
			if err != nil {
				return result, duplicates, fmt.Errorf("cannot determine link targets in document %q: %w", docs[i].FileName(), err)
			}
			err = result.Add(linkTarget)
			if err != nil {
				duplicates = append(duplicates, linkTarget)
			}
		}
	}

	// add link targets for resources
	for _, fileName := range resources.FileNames() {
		err := result.Add(fileName)
		if err != nil {
			duplicates = append(duplicates, fileName)
		}
	}
	return result, duplicates, nil
}

func (ltc linkTargets) Add(linkTarget string) error {
	if ltc.Contains(linkTarget) {
		return fmt.Errorf("duplicate link target: %s", linkTarget)
	}
	ltc[linkTarget] = struct{}{}
	return nil
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
