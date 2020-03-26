package check

import (
	"strings"

	"github.com/kevgo/tikibase/domain"
)

// docsWithEmptySections provides the filenames of documents that contain empty sections.
func docsWithEmptySections(docs domain.Documents) (result []string, err error) {
	for d := range docs {
		sections := docs[d].AllSections()
		for s := range sections {
			title, err := sections[s].Title()
			if err != nil {
				return result, err
			}
			if strings.TrimSpace(title) == "" {
				result = append(result, docs[d].FileName())
			}
		}
	}
	return result, nil
}
