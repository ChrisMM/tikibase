package mentions

import (
	"fmt"

	"github.com/kevgo/tikibase/domain"
)

// RenderMentionsSection provides a domain.Section containing the given mentions of a document.
func RenderMentionsSection(links domain.TikiLinkCollection, doc *domain.Document) domain.Section {
	builder := domain.NewSectionBuilder("### mentions\n\n", doc)
	for i := range links {
		sourceDoc := links[i].SourceSection().Document()
		builder.AddLine(fmt.Sprintf("- [%s](%s)\n", sourceDoc.TitleSection().Title(), links[i].SourceSection().URL()))
	}
	return builder.Section()
}
