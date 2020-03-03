package occurrences

import (
	"fmt"

	"github.com/kevgo/tikibase/domain"
)

// CreateOccurrencesSection provides a domain.Section containing the given occurrences of a document.
func CreateOccurrencesSection(links domain.TikiLinkCollection, doc *domain.Document) (result *domain.Section, err error) {
	builder := domain.NewSectionBuilder("### occurrences\n\n", doc)
	for i := range links {
		sourceSection := links[i].SourceSection()
		sourceDoc := sourceSection.Document()
		sourceDocTitleSection := sourceDoc.TitleSection()
		sourceDocTitleSectionTitle, err := sourceDocTitleSection.Title()
		if err != nil {
			return result, err
		}
		if sourceSection == sourceDocTitleSection {
			// the link occurs in the title section of the document --> link to the document
			builder.AddLine(fmt.Sprintf("- [%s](%s)\n", sourceDocTitleSectionTitle, sourceDoc.URL()))
		} else {
			// the link occurs in a content section of the document --> link to the section in the document
			sourceSectionTitle, err := sourceSection.Title()
			if err != nil {
				return result, err
			}
			sourceSectionURL, err := sourceSection.URL()
			if err != nil {
				return result, err
			}
			builder.AddLine(fmt.Sprintf("- [%s (%s)](%s)\n", sourceDocTitleSectionTitle, sourceSectionTitle, sourceSectionURL))
		}
	}
	return builder.Section(), nil
}
