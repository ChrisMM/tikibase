package domain_test

import (
	"testing"

	"github.com/kevgo/tikibase/domain"
	"github.com/stretchr/testify/assert"
)

func TestSectionBuilder(t *testing.T) {
	doc := domain.ScaffoldDocument(domain.DocumentScaffold{})
	tsb := domain.NewSectionBuilder("# Title\n", doc)
	tsb.AddLine("\n")
	tsb.AddLine("content 1\n")
	tsb.AddLine("content 2\n")
	tsb.AddLine("\n")
	section := tsb.Section()
	assert.Equal(t, domain.SectionContent("# Title\n\ncontent 1\ncontent 2\n\n"), section.Content(), "TikiSectionBuilder didn't build the right content")
	assert.Same(t, doc, section.Document(), "Created section doesn't contain a link to its containing document")
}
