package domain_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kevgo/tikibase/domain"
)

func TestTikiDocumentCollectionFileNames(t *testing.T) {
	docs := domain.ScaffoldTikiDocumentCollection([]domain.TikiDocumentScaffold{
		{FileName: "one.md"},
		{FileName: "two.md"},
	})
	actual, err := docs.FileNames()
	if err != nil {
		t.Fatalf("cannot get filenames of docs: %v", err)
	}
	expected := []domain.TikiDocumentFilename{
		domain.TikiDocumentFilename("one.md"),
		domain.TikiDocumentFilename("two.md"),
	}
	diff := cmp.Diff(expected, actual)
	if diff != "" {
		t.Fatal(diff)
	}
}

func TestTikiDocumentCollectionTikiLinks(t *testing.T) {
	docs := domain.ScaffoldTikiDocumentCollection([]domain.TikiDocumentScaffold{
		{FileName: "one.md", Content: "# The one\n[The other](two.md)"},
		{FileName: "two.md", Content: "# The other\n[The one](one.md)"},
	})
	actual, err := docs.TikiLinks()
	if err != nil {
		t.Fatalf("cannot get TikiLinks of docs: %v", err)
	}
	expected := domain.TikiLinkCollection{
		domain.NewTikiLink("The other", docs[0].TitleSection(), docs[1]),
		domain.NewTikiLink("The one", docs[1].TitleSection(), docs[0]),
	}
	diff := cmp.Diff(expected, actual, cmp.AllowUnexported(expected[0], expected[0].SourceSection(), expected[0].TargetDocument()))
	if diff != "" {
		t.Fatal(diff)
	}
}
