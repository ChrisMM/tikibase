package mentions

import (
	"fmt"

	"github.com/kevgo/tikibase/domain"
)

// RenderMentionsSection provides a domain.Section containing the given mentions of a document.
func RenderMentionsSection(links domain.TikiLinkCollection, doc *domain.Document) (result domain.Section, err error) {
	builder := domain.NewSectionBuilder("### mentions\n\n", doc)
	for i := range links {
		sourceSection := links[i].SourceSection()
		sourceDoc := sourceSection.Document()
		sourceDocTitleSectionTitle, err := sourceDoc.TitleSection().Title()
		if err != nil {
			return result, err
		}
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
	return builder.Section(), nil
}
