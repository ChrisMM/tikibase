package domain_test

import (
	"testing"

	"github.com/kevgo/tikibase/domain"
	"github.com/stretchr/testify/assert"
)

func TestDocumentAllSections(t *testing.T) {
	doc := domain.ScaffoldDocument(domain.DocumentScaffold{
		Content: "# Title\n\nmy doc\n\n### One\n\nThe one.\n\n### Two\n\nThe other.\n",
	})
	sections := doc.AllSections()
	assert.Len(t, sections, 3)
	assert.Equal(t, domain.SectionContent("# Title\n\nmy doc\n\n"), (sections)[0].Content())
	assert.Equal(t, domain.SectionContent("### One\n\nThe one.\n\n"), (sections)[1].Content())
	assert.Equal(t, domain.SectionContent("### Two\n\nThe other.\n"), (sections)[2].Content())
}

func TestDocumentAppendSection(t *testing.T) {
	oldDoc := domain.ScaffoldDocument(domain.DocumentScaffold{FileName: "one.md", Content: "existing document content\n"})
	newSection := domain.ScaffoldSection(domain.SectionScaffold{Content: "### new section\n"})
	newDoc := oldDoc.AppendSection(newSection)
	assert.Equal(t, "existing document content\n\n### new section\n", newDoc.Content())
	assert.Equal(t, "one.md", newDoc.FileName(), "didn't bring the filename over to the new doc")
}

func TestDocumentContentSections(t *testing.T) {
	doc := domain.ScaffoldDocument(domain.DocumentScaffold{
		Content: "# Title\nmy doc\n### One\nThe one.\n### Two\nThe other.\n",
	})
	sections := doc.ContentSections()
	assert.Len(t, sections, 2, "unexpected sections length")
	assert.Equal(t, domain.SectionContent("### One\nThe one.\n"), (sections)[0].Content())
	assert.Equal(t, domain.SectionContent("### Two\nThe other.\n"), (sections)[1].Content())
}

func TestDocumentFileName(t *testing.T) {
	doc := domain.ScaffoldDocument(domain.DocumentScaffold{FileName: "one.md"})
	assert.Equal(t, "one.md", doc.FileName())
}

func TestDocumentId(t *testing.T) {
	tests := map[string]string{
		"foo.md":         "foo",
		"markdown-it.md": "markdown-it",
	}
	for give, want := range tests {
		t.Run(give, func(t *testing.T) {
			doc := domain.ScaffoldDocument(domain.DocumentScaffold{FileName: give})
			assert.Equal(t, want, doc.ID())
		})
	}
}

func TestDocumentReplaceSection(t *testing.T) {
	td := domain.ScaffoldDocument(domain.DocumentScaffold{
		Content: "# Title\n\nmy doc\n\n### One\n\nThe one.\n\n### Two\n\nOld section 2.\n",
	})
	sections := td.AllSections()
	twoSection := (sections)[2]
	newSection := domain.ScaffoldSection(domain.SectionScaffold{Content: "### Two\n\nNew section 2.\n", Doc: td})
	newdoc := td.ReplaceSection(twoSection, newSection)
	newSections := newdoc.AllSections()
	assert.Len(t, newSections, 3, "unexpected newSections length")
	assert.Equal(t, domain.SectionContent("# Title\n\nmy doc\n\n"), (newSections)[0].Content(), "unexpected title section")
	assert.Equal(t, domain.SectionContent("### One\n\nThe one.\n\n"), (newSections)[1].Content(), "unexpected content section 1")
	assert.Equal(t, domain.SectionContent("### Two\n\nNew section 2.\n"), (newSections)[2].Content(), "unexpected content section 2")
}

func TestDocumentTikiLinks(t *testing.T) {
	docs := domain.ScaffoldDocumentCollection([]domain.DocumentScaffold{
		{FileName: "doc1.md", Content: "### One\n"},
		{FileName: "doc2.md", Content: "### Two\n[one](doc1.md)"},
	})
	actual, err := docs[1].TikiLinks(docs)
	assert.Nil(t, err, "error getting TikiLinks for doc2")
	assert.Len(t, actual, 1)
	assert.Equal(t, "one", actual[0].Title())
	assert.Equal(t, docs[1].TitleSection(), actual[0].SourceSection())
	assert.Equal(t, docs[0], actual[0].TargetDocument())
}

func TestDocumentTitleSection(t *testing.T) {
	doc := domain.ScaffoldDocument(domain.DocumentScaffold{Content: "# My Title\n\nTitle section content.\n\n### Content Section\n Content section content.\n"})
	section := doc.TitleSection()
	assert.Equal(t, domain.SectionContent("# My Title\n\nTitle section content.\n\n"), section.Content())
}

func TestScaffoldDocument(t *testing.T) {
	actual := domain.ScaffoldDocument(domain.DocumentScaffold{})
	assert.NotEqual(t, "", actual.FileName(), "no default FileName")
	titleSectionTitle, err := actual.TitleSection().Title()
	assert.Nil(t, err)
	assert.NotEqual(t, "", titleSectionTitle, "no default section")
}
