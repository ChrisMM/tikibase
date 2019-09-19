package domain_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kevgo/tikibase/domain"
)

func TestTikiLinkCollectionEqual(t *testing.T) {
	docs := domain.ScaffoldDocumentCollection([]domain.DocumentScaffold{{}, {}})
	expected := domain.ScaffoldTikiLinkCollection([]domain.TikiLinkScaffold{
		{Title: "1-2", SourceSection: docs[0].TitleSection(), TargetDocument: &docs[1]},
		{Title: "2-1", SourceSection: docs[1].TitleSection(), TargetDocument: &docs[0]},
	})

	// compare against a TikiLinkCollection with similar contents
	equal := domain.ScaffoldTikiLinkCollection([]domain.TikiLinkScaffold{
		{Title: "1-2", SourceSection: docs[0].TitleSection(), TargetDocument: &docs[1]},
		{Title: "2-1", SourceSection: docs[1].TitleSection(), TargetDocument: &docs[0]},
	})
	if diff := cmp.Diff(expected, equal); diff != "" {
		t.Fatalf("equal: didn't match: %s", diff)
	}

	// compare against a longer TikiLinkCollection
	longer := domain.ScaffoldTikiLinkCollection([]domain.TikiLinkScaffold{
		{Title: "1-2", SourceSection: docs[0].TitleSection(), TargetDocument: &docs[1]},
		{Title: "2-1", SourceSection: docs[1].TitleSection(), TargetDocument: &docs[0]},
		{Title: "1-2", SourceSection: docs[0].TitleSection(), TargetDocument: &docs[1]},
	})
	if diff := cmp.Diff(expected, longer); diff == "" {
		t.Fatalf("longer: unexpected match: %v", diff)
	}

	// compare against a shorter TikiLinkCollection
	shorter := domain.ScaffoldTikiLinkCollection([]domain.TikiLinkScaffold{
		{Title: "1-2", SourceSection: docs[0].TitleSection(), TargetDocument: &docs[1]},
	})
	if diff := cmp.Diff(expected, shorter); diff == "" {
		t.Fatalf("shorter: unexpected match: %v", diff)
	}

	// compare against a TikiLinkCollection with a different section
	differentSection := domain.ScaffoldTikiLinkCollection([]domain.TikiLinkScaffold{
		{Title: "1-2", SourceSection: docs[0].TitleSection(), TargetDocument: &docs[1]},
		{Title: "2-1", SourceSection: docs[0].TitleSection(), TargetDocument: &docs[0]},
	})
	if diff := cmp.Diff(expected, differentSection); diff == "" {
		t.Fatalf("differentSection: unexpected match: %s", diff)
	}

	// compare against a TikiLinkCollection with a different target document
	differentDoc := domain.ScaffoldTikiLinkCollection([]domain.TikiLinkScaffold{
		{Title: "1-2", SourceSection: docs[0].TitleSection(), TargetDocument: &docs[1]},
		{Title: "2-1", SourceSection: docs[1].TitleSection(), TargetDocument: &docs[1]},
	})
	if diff := cmp.Diff(expected, differentDoc); diff == "" {
		t.Fatalf("differentDoc: unexpected match: %s", diff)
	}
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
		{Title: "1-2", SourceSection: docs[0].TitleSection(), TargetDocument: &docs[1]},
		{Title: "1-3", SourceSection: docs[0].TitleSection(), TargetDocument: &docs[2]},
		{Title: "2-1", SourceSection: docs[1].TitleSection(), TargetDocument: &docs[0]},
		{Title: "2-3", SourceSection: docs[1].TitleSection(), TargetDocument: &docs[2]},
		{Title: "3-1", SourceSection: docs[2].TitleSection(), TargetDocument: &docs[0]},
		{Title: "3-2", SourceSection: docs[2].TitleSection(), TargetDocument: &docs[1]},
	})
	actual := links.GroupByTarget()

	// verify
	expected := map[domain.DocumentFilename]domain.TikiLinkCollection{
		domain.DocumentFilename("one.md"):   {links[2], links[4]},
		domain.DocumentFilename("two.md"):   {links[0], links[5]},
		domain.DocumentFilename("three.md"): {links[1], links[3]},
	}
	if diff := cmp.Diff(expected, actual); diff != "" {
		t.Fatal(diff)
	}
}

func TestTikiLinkCollectionScaffold(t *testing.T) {
	actual := domain.ScaffoldTikiLinkCollection([]domain.TikiLinkScaffold{
		{Title: "foo"},
	})
	if actual[0].Title() != "foo" {
		t.Fatal("didn't scaffold a TikiLink")
	}
}
