package domain_test

import (
	"testing"

	"github.com/kevgo/tikibase/domain"
	"github.com/stretchr/testify/assert"
)

func TestDocuments_Contains(t *testing.T) {
	docs := domain.ScaffoldDocuments([]domain.DocumentScaffold{{}})
	otherDoc := domain.ScaffoldDocument(domain.DocumentScaffold{})
	assert.True(t, docs.Contains(docs[0]), "looking contain existing doc")
	assert.False(t, docs.Contains(otherDoc), "should not contain otherDoc")
}

func TestDocuments_Find(t *testing.T) {
	docs := domain.ScaffoldDocuments([]domain.DocumentScaffold{
		{FileName: "one.md"},
		{FileName: "two.md"},
	})
	actual, err := docs.FindByFilename("two.md")
	assert.Nil(t, err, "cannot find document 'two.md'")
	assert.Equal(t, "two.md", actual.FileName())
}

func TestDocuments_Links(t *testing.T) {
	docs := domain.ScaffoldDocuments([]domain.DocumentScaffold{
		{FileName: "one.md", Content: "# The one\n[The other](two.md)"},
		{FileName: "two.md", Content: "# The other\n[Google](http://google.com)"},
	})
	internalLinks, externalLinks := docs.Links()
	assert.Len(t, internalLinks, 1)
	assert.Equal(t, internalLinks[0].Title(), "The other")
	assert.Same(t, internalLinks[0].SourceSection(), docs[0].TitleSection())
	assert.Equal(t, internalLinks[0].Target(), "two.md")
	assert.Len(t, externalLinks, 1)
	assert.Equal(t, externalLinks[0].Title(), "Google")
	assert.Same(t, externalLinks[0].SourceSection(), docs[1].TitleSection())
	assert.Equal(t, externalLinks[0].Target(), "http://google.com")
}

func TestDocuments_TikiLinks(t *testing.T) {
	docs := domain.ScaffoldDocuments([]domain.DocumentScaffold{
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
