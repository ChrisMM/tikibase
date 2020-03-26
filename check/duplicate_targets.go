package check

import (
	"fmt"

	"github.com/kevgo/tikibase/domain"
)

func findDuplicateTargets(docs domain.Documents) (result []string, err error) {
	knownTargets := make(map[string]struct{})
	for d := range docs {
		sections := docs[d].AllSections()
		for s := range sections {
			linkTarget, err := sections[s].LinkTarget()
			if err != nil {
				return result, fmt.Errorf("cannot determine link targets in document %q: %w", docs[d].FileName(), err)
			}
			if _, has := knownTargets[linkTarget]; has {
				result = append(result, linkTarget)
			}
			knownTargets[linkTarget] = struct{}{}
		}
	}
	return result, nil
}
