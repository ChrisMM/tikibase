package domain_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kevgo/tikibase/domain"
)

func TestDocumentCollectionFileNames(t *testing.T) {
	docs := domain.ScaffoldDocumentCollection([]domain.DocumentScaffold{
		{FileName: "one.md"},
		{FileName: "two.md"},
	})
	actual, err := docs.FileNames()
	if err != nil {
		t.Fatalf("cannot get filenames of docs: %v", err)
	}
	expected := []domain.DocumentFilename{
		domain.DocumentFilename("one.md"),
		domain.DocumentFilename("two.md"),
	}
	diff := cmp.Diff(expected, actual)
	if diff != "" {
		t.Fatal(diff)
	}
}

func TestDocumentCollectionTikiLinks(t *testing.T) {
	docs := domain.ScaffoldDocumentCollection([]domain.DocumentScaffold{
		{FileName: "one.md", Content: "# The one\n[The other](two.md)"},
		{FileName: "two.md", Content: "# The other\n[The one](one.md)"},
	})
	actual, err := docs.TikiLinks()
	if err != nil {
		t.Fatalf("cannot get TikiLinks of docs: %v", err)
	}
	expected := domain.ScaffoldTikiLinkCollection([]domain.TikiLinkScaffold{
		{Title: "The other", SourceSection: docs[0].TitleSection(), TargetDocument: docs[1]},
		{Title: "The one", SourceSection: docs[1].TitleSection(), TargetDocument: docs[0]},
	})
	diff := cmp.Diff(expected, actual, cmp.AllowUnexported(expected[0], expected[0].SourceSection(), expected[0].TargetDocument()))
	if diff != "" {
		t.Fatal(diff)
	}
}