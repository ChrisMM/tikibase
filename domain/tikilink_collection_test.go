package domain_test

import (
	"testing"

	"github.com/kevgo/tikibase/domain"
	"github.com/stretchr/testify/assert"
)

func TestTikiLinkCollectionContains(t *testing.T) {
	docs := domain.ScaffoldDocumentCollection([]domain.DocumentScaffold{{}, {}})
	links := domain.ScaffoldTikiLinkCollection([]domain.TikiLinkScaffold{
		{Title: "1-2", SourceSection: docs[0].TitleSection(), TargetDocument: docs[1]},
	})
	containedLink := domain.ScaffoldTikiLink(domain.TikiLinkScaffold{Title: "1-2", SourceSection: docs[0].TitleSection(), TargetDocument: docs[1]})
	notContainedLink := domain.ScaffoldTikiLink(domain.TikiLinkScaffold{Title: "2-1", SourceSection: docs[1].TitleSection(), TargetDocument: docs[0]})
	assert.True(t, links.Contains(containedLink), "expected collection to contain this link")
	assert.False(t, links.Contains(notContainedLink), "expected collection to not contain this link")
}

func TestTikiLinkCollectionGroupByTarget(t *testing.T) {
	// create documents
	docs := domain.ScaffoldDocumentCollection([]domain.DocumentScaffold{
		{FileName: "one.md"},
		{FileName: "two.md"},
		{FileName: "three.md"},
	})

	// convert links
	links := domain.ScaffoldTikiLinkCollection([]domain.TikiLinkScaffold{
		{Title: "1-2", SourceSection: docs[0].TitleSection(), TargetDocument: docs[1]},
		{Title: "1-3", SourceSection: docs[0].TitleSection(), TargetDocument: docs[2]},
		{Title: "2-1", SourceSection: docs[1].TitleSection(), TargetDocument: docs[0]},
		{Title: "2-3", SourceSection: docs[1].TitleSection(), TargetDocument: docs[2]},
		{Title: "3-1", SourceSection: docs[2].TitleSection(), TargetDocument: docs[0]},
		{Title: "3-2", SourceSection: docs[2].TitleSection(), TargetDocument: docs[1]},
	})
	actual := links.GroupByTarget()

	// verify
	assert.Equal(t, actual[domain.DocumentFilename("one.md")], domain.TikiLinkCollection{links[2], links[4]})
	assert.Equal(t, actual[domain.DocumentFilename("two.md")], domain.TikiLinkCollection{links[0], links[5]})
	assert.Equal(t, actual[domain.DocumentFilename("three.md")], domain.TikiLinkCollection{links[1], links[3]})
}

func TestTikiLinkCollectionScaffold(t *testing.T) {
	actual := domain.ScaffoldTikiLinkCollection([]domain.TikiLinkScaffold{
		{Title: "foo"},
	})
	assert.Equal(t, "foo", actual[0].Title(), "didn't scaffold a TikiLink")
}

func TestTikiLinkCollectionUnique(t *testing.T) {
	docs := domain.ScaffoldDocumentCollection([]domain.DocumentScaffold{
		{FileName: "one.md"},
		{FileName: "two.md"},
	})
	links := domain.ScaffoldTikiLinkCollection([]domain.TikiLinkScaffold{
		{Title: "1-2", SourceSection: docs[0].TitleSection(), TargetDocument: docs[1]},
		{Title: "1-2", SourceSection: docs[0].TitleSection(), TargetDocument: docs[1]},
	})
	assert.Len(t, links.Unique(), 1)
}
