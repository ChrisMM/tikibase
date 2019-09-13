package domain_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kevgo/tikibase/domain"
)

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
	expected := map[domain.DocumentFilename]domain.TikiLinkCollection{
		domain.DocumentFilename("one.md"):   {links[2], links[4]},
		domain.DocumentFilename("two.md"):   {links[0], links[5]},
		domain.DocumentFilename("three.md"): {links[1], links[3]},
	}
	diff := cmp.Diff(expected, actual, cmp.AllowUnexported(links[0], docs[0].TitleSection(), docs[0]))
	if diff != "" {
		t.Fatal(diff)
	}
}
