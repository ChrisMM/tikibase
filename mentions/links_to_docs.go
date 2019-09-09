package mentions

import (
	"github.com/kevgo/tikibase/domain"
)

// LinksToDocs determines which links in the given TikiLink list point to which TikiDocument.
func LinksToDocs(links []domain.TikiLink) map[domain.TikiDocumentFilename][]domain.TikiLink {
	result := map[domain.TikiDocumentFilename][]domain.TikiLink{}
	for _, link := range links {
		doc := link.TargetDocument()
		targetFileName := doc.FileName()
		result[targetFileName] = append(result[targetFileName], link)
	}
	return result
}
