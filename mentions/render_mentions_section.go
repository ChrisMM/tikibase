package mentions

import (
	"fmt"

	"github.com/kevgo/tikibase/domain"
)

// RenderMentionsSection provides a domain.Section containing the given mentions of a document.
func RenderMentionsSection(links domain.TikiLinkCollection, doc *domain.Document) domain.Section {
	builder := domain.NewSectionBuilder("### mentions\n\n", doc)
	for _, link := range links {
		sourceDoc := link.SourceSection().Document()
		builder.AddLine(fmt.Sprintf("- [%s](%s)\n", sourceDoc.TitleSection().Title(), link.SourceSection().URL()))
	}
	return builder.Section()
}
