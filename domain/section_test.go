package domain_test

import (
	"testing"

	"github.com/kevgo/tikibase/domain"
	"github.com/stretchr/testify/assert"
)

func TestSection_Anchor(t *testing.T) {
	section := domain.ScaffoldSection(domain.SectionScaffold{Content: "### what is it\n"})
	actual, err := section.Anchor()
	assert.Nil(t, err)
	assert.Equal(t, "what-is-it", actual)
}

func TestSection_AppendLine(t *testing.T) {
	section := domain.ScaffoldSection(domain.SectionScaffold{Content: "existing content\n"})
	newSection := section.AppendText("new line\n")
	assert.Equal(t, domain.SectionContent("existing content\nnew line\n"), newSection.Content())
}

func TestSection_Content(t *testing.T) {
	expectedContent := "### title\nthe content\n"
	section := domain.ScaffoldSection(domain.SectionScaffold{Content: expectedContent})
	assert.Equal(t, domain.SectionContent(expectedContent), section.Content())
}

func TestSection_LinkTarget(t *testing.T) {
	doc := domain.ScaffoldDocument(domain.DocumentScaffold{FileName: "doc.md"})
	section := domain.ScaffoldSection(domain.SectionScaffold{"### Section 3", doc})
	actual, err := section.LinkTarget()
	assert.Nil(t, err)
	assert.Equal(t, "doc.md#section-3", actual)
}

func TestSection_TikiLinks(t *testing.T) {
	docs := domain.ScaffoldDocumentCollection([]domain.DocumentScaffold{
		{FileName: "one.md", Content: "# One"},
		{FileName: "two.md", Content: `# Title\ntext [MD link to doc1](one.md)\n text [MD link to doc2](two.md) text\ntext <a href="one.md">HTML link to doc1</a> text <a textrun="dope">not a link</a>`},
	})
	section := docs[1].TitleSection()
	actual, err := section.TikiLinks(docs)
	assert.Nil(t, err, "cannot get links in section")
	assert.Len(t, actual, 3)
	assert.Equal(t, "MD link to doc1", actual[0].Title())
	assert.Same(t, section, actual[0].SourceSection())
	assert.Same(t, docs[0], actual[0].TargetDocument())
	assert.Equal(t, "MD link to doc2", actual[1].Title())
	assert.Same(t, section, actual[1].SourceSection())
	assert.Same(t, docs[1], actual[1].TargetDocument())
	assert.Equal(t, "HTML link to doc1", actual[2].Title())
	assert.Same(t, section, actual[2].SourceSection())
	assert.Same(t, docs[0], actual[2].TargetDocument())
}

func TestSection_TikiLinks_IgnoresHtmlLinks(t *testing.T) {
	docs := domain.ScaffoldDocumentCollection([]domain.DocumentScaffold{
		{Content: "# Title\ntext [HTML link](http://google.com)"},
	})
	actual, err := docs[0].TitleSection().TikiLinks(docs)
	assert.Nil(t, err)
	assert.Len(t, actual, 0)
}

func TestSection_Title(t *testing.T) {
	section := domain.ScaffoldSection(domain.SectionScaffold{Content: "### What is it\n"})
	actual, err := section.Title()
	assert.Nil(t, err)
	assert.Equal(t, "What is it", actual, "mismatching section title")
}

func TestSection_URL(t *testing.T) {
	doc := domain.ScaffoldDocument(domain.DocumentScaffold{FileName: "one.md"})
	section := domain.ScaffoldSection(domain.SectionScaffold{Content: "### What is it\n", Doc: doc})
	actual, err := section.URL()
	assert.Nil(t, err)
	assert.Equal(t, "one.md#what-is-it", actual)
}
