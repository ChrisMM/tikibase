package domain_test

import (
	"testing"

	"github.com/kevgo/tikibase/domain"
	"github.com/stretchr/testify/assert"
)

func TestTikiLinks_CombineLinksFromSameDocuments(t *testing.T) {
	docs := domain.ScaffoldDocuments([]domain.DocumentScaffold{
		{Content: "# One\n\n### sec1\n\n### sec2\n"},
		{},
	})
	doc0sections := docs[0].AllSections()
	links := domain.ScaffoldTikiLinks([]domain.TikiLinkScaffold{
		{Title: "0a-1", SourceSection: (doc0sections)[0], TargetDocument: docs[1]},
		{Title: "0b-1", SourceSection: (doc0sections)[1], TargetDocument: docs[1]},
	})
	actual := links.CombineLinksFromSameDocuments()
	assert.Len(t, actual, 1)
	assert.Same(t, docs[0].TitleSection(), actual[0].SourceSection())
	assert.Same(t, docs[1], actual[0].TargetDocument())
}

func TestTikiLinks_Contains(t *testing.T) {
	docs := domain.ScaffoldDocuments([]domain.DocumentScaffold{{}, {}})
	links := domain.ScaffoldTikiLinks([]domain.TikiLinkScaffold{
		{Title: "0-1", SourceSection: docs[0].TitleSection(), TargetDocument: docs[1]},
	})
	containedLink := links[0]
	notContainedLink := domain.ScaffoldTikiLink(domain.TikiLinkScaffold{Title: "2-1", SourceSection: docs[1].TitleSection(), TargetDocument: docs[0]})
	assert.True(t, links.Contains(containedLink))
	assert.False(t, links.Contains(notContainedLink))
}

func TestTikiLinks_Filter(t *testing.T) {
	links := domain.ScaffoldTikiLinks([]domain.TikiLinkScaffold{{}, {}})
	actual := links.Filter(func(link *domain.TikiLink) bool { return true })
	assert.Len(t, actual, 2)
	actual = links.Filter(func(link *domain.TikiLink) bool { return false })
	assert.Len(t, actual, 0)
}

func TestTikiLinks_GroupByTarget(t *testing.T) {
	docs := domain.ScaffoldDocuments([]domain.DocumentScaffold{
		{FileName: "0.md"},
		{FileName: "1.md"},
		{FileName: "2.md"},
	})
	links := domain.ScaffoldTikiLinks([]domain.TikiLinkScaffold{
		{Title: "0-1", SourceSection: docs[0].TitleSection(), TargetDocument: docs[1]},
		{Title: "0-2", SourceSection: docs[0].TitleSection(), TargetDocument: docs[2]},
		{Title: "1-0", SourceSection: docs[1].TitleSection(), TargetDocument: docs[0]},
		{Title: "1-2", SourceSection: docs[1].TitleSection(), TargetDocument: docs[2]},
		{Title: "2-0", SourceSection: docs[2].TitleSection(), TargetDocument: docs[0]},
		{Title: "2-1", SourceSection: docs[2].TitleSection(), TargetDocument: docs[1]},
	})
	actual := links.GroupByTarget()
	assert.Equal(t, actual["0.md"], domain.TikiLinks{links[2], links[4]})
	assert.Equal(t, actual["1.md"], domain.TikiLinks{links[0], links[5]})
	assert.Equal(t, actual["2.md"], domain.TikiLinks{links[1], links[3]})
}

func TestTikiLinks_ReferencedDocs(t *testing.T) {
	docs := domain.ScaffoldDocuments([]domain.DocumentScaffold{{}, {}, {}})
	links := domain.ScaffoldTikiLinks([]domain.TikiLinkScaffold{
		{Title: "0-1", SourceSection: docs[0].TitleSection(), TargetDocument: docs[1]},
		{Title: "1-0", SourceSection: docs[1].TitleSection(), TargetDocument: docs[0]},
		{Title: "2-0", SourceSection: docs[2].TitleSection(), TargetDocument: docs[0]},
	})
	actual := links.ReferencedDocs()
	assert.Len(t, actual, 2)
	assert.Same(t, docs[1], actual[0], "actual[0] should == docs[1]")
	assert.Same(t, docs[0], actual[1], "actual[1] should == docs[0]")
}

func TestTikiLinks_RemoveLinksFromDocs(t *testing.T) {
	docs := domain.ScaffoldDocuments([]domain.DocumentScaffold{{}, {}})
	doc2 := domain.ScaffoldDocument(domain.DocumentScaffold{FileName: "2.md"})
	links := domain.ScaffoldTikiLinks([]domain.TikiLinkScaffold{
		{Title: "0-1", SourceSection: docs[0].TitleSection(), TargetDocument: docs[1]},
		{Title: "1-0", SourceSection: docs[1].TitleSection(), TargetDocument: docs[0]},
		{Title: "2-0", SourceSection: doc2.TitleSection(), TargetDocument: docs[0]},
	})
	newLinks := links.RemoveLinksFromDocs(docs)
	assert.Len(t, newLinks, 1)
	assert.Same(t, links[2], newLinks[0], "should contain the last link")
}

func TestTikiLinks_Scaffold(t *testing.T) {
	actual := domain.ScaffoldTikiLinks([]domain.TikiLinkScaffold{
		{Title: "foo"},
	})
	assert.Equal(t, "foo", actual[0].Title())
}

func TestTikiLinks_SortBySourceDocumentTitle(t *testing.T) {
	docs := domain.ScaffoldDocuments([]domain.DocumentScaffold{
		{Content: "# One"},
		{Content: "# Two"},
		{Content: "# Three"},
	})
	links := domain.ScaffoldTikiLinks([]domain.TikiLinkScaffold{
		{SourceSection: docs[0].TitleSection()},
		{SourceSection: docs[1].TitleSection()},
		{SourceSection: docs[2].TitleSection()},
	})
	links.SortBySourceDocumentTitle()
	assert.Len(t, links, 3)
	assert.Same(t, docs[0].TitleSection(), links[0].SourceSection())
	assert.Same(t, docs[2].TitleSection(), links[1].SourceSection())
	assert.Same(t, docs[1].TitleSection(), links[2].SourceSection())
}

func TestTikiLinks_Unique(t *testing.T) {
	docs := domain.ScaffoldDocuments([]domain.DocumentScaffold{
		{FileName: "0.md"},
		{FileName: "1.md"},
	})
	links := domain.ScaffoldTikiLinks([]domain.TikiLinkScaffold{
		{Title: "0-1", SourceSection: docs[0].TitleSection(), TargetDocument: docs[1]},
		{Title: "0-1", SourceSection: docs[0].TitleSection(), TargetDocument: docs[1]},
	})
	assert.Len(t, links.Unique(), 1)
}
