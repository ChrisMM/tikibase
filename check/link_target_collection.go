package check

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kevgo/tikibase/domain"
)

// linkTargetCollection contains all possible link targets within this Tikibase.
//
// Examples: `1.md`, `1.md#foo`
type linkTargetCollection map[string]struct{}

func findLinkTargets(fileNames []string, docs domain.DocumentCollection) (result linkTargetCollection, duplicates []string, err error) {
	result = make(linkTargetCollection)

	// add links targets for documents
	for i := range docs {

		// add target for the document itself
		err := result.Add(string(docs[i].FileName()))
		if err != nil {
			duplicates = append(duplicates, string(docs[i].FileName()))
		}

		// add target for the sections in the document
		sections := docs[i].AllSections()
		for k := range sections {
			linkTarget, err := sections[k].LinkTarget()
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
	for i := range fileNames {
		if !strings.HasSuffix(fileNames[i], ".md") {
			err := result.Add(fileNames[i])
			if err != nil {
				duplicates = append(duplicates, fileNames[i])
			}
		}
	}
	return result, duplicates, nil
}

func (ltc linkTargetCollection) Add(linkTarget string) error {
	if ltc.Contains(linkTarget) {
		return fmt.Errorf("duplicate link target: %s", linkTarget)
	}
	ltc[linkTarget] = struct{}{}
	return nil
}

func (ltc linkTargetCollection) Contains(linkTarget string) bool {
	_, ok := ltc[linkTarget]
	return ok
}

func (ltc linkTargetCollection) String() (result string) {
	values := []string{}
	for value := range ltc {
		values = append(values, value)
	}
	sort.Strings(values)
	return "[" + strings.Join(values, ", ") + "]"
}
