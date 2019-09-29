package domain_test

import (
	"testing"

	"github.com/kevgo/tikibase/domain"
	"github.com/stretchr/testify/assert"
)

func TestDocumentCollectionContains(t *testing.T) {
	docs := domain.ScaffoldDocumentCollection([]domain.DocumentScaffold{
		{FileName: "one.md"},
		{FileName: "two.md"},
	})
	otherDoc := domain.ScaffoldDocument(domain.DocumentScaffold{})
	assert.True(t, docs.Contains(docs[0]), "looking contain existing doc")
	assert.False(t, docs.Contains(otherDoc), "should not contain otherDoc")
}

func TestDocumentCollectionFileNames(t *testing.T) {
	docs := domain.ScaffoldDocumentCollection([]domain.DocumentScaffold{
		{FileName: "one.md"},
		{FileName: "two.md"},
	})
	result, err := docs.FileNames()
	assert.Nil(t, err, "cannot get filenames of docs")
	assert.Equal(t, len(result), 2)
	assert.Equal(t, result[0], domain.DocumentFilename("one.md"))
	assert.Equal(t, result[1], domain.DocumentFilename("two.md"))
}

func TestDocumentCollectionFind(t *testing.T) {
	docs := domain.ScaffoldDocumentCollection([]domain.DocumentScaffold{
		{FileName: "one.md", Content: "# The one\n[The other](two.md)"},
		{FileName: "two.md", Content: "# The other\n[The one](one.md)"},
	})
	actual, err := docs.Find("two.md")
	assert.Nil(t, err, "cannot find document 'two.md'")
	assert.Equal(t, domain.DocumentFilename("two.md"), actual.FileName(), "found the wrong document")
}

func TestDocumentCollectionTikiLinks(t *testing.T) {
	docs := domain.ScaffoldDocumentCollection([]domain.DocumentScaffold{
		{FileName: "one.md", Content: "# The one\n[The other](two.md)"},
		{FileName: "two.md", Content: "# The other\n[The one](one.md)"},
	})
	actual, err := docs.TikiLinks()
	assert.Nil(t, err, "cannot get TikiLinks of docs")
	assert.Len(t, actual, 2)
	assert.Equal(t, actual[0].Title(), "The other")
	assert.Same(t, actual[0].SourceSection(), docs[0].TitleSection())
	assert.Same(t, actual[0].TargetDocument(), docs[1])
	assert.Equal(t, actual[1].Title(), "The one")
	assert.Same(t, actual[1].SourceSection(), docs[1].TitleSection())
	assert.Same(t, actual[1].TargetDocument(), docs[0])
}
