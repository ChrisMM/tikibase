package check

import "github.com/kevgo/tikibase/domain"

func findNonLinkedResources(resourceFiles domain.ResourceFiles, internalLinks domain.Links) (result []string) {
	for r := range resourceFiles {
		if !internalLinks.HasLinkTo(resourceFiles[r]) {
			result = append(result, resourceFiles[r])
		}
	}
	return result
}
