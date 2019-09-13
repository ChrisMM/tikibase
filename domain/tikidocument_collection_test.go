package domain_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kevgo/tikibase/domain"
	"github.com/kevgo/tikibase/test"
)

func TestTikiDocumentCollectionFileNames(t *testing.T) {
	tb, dc := test.NewDocumentCreator(t)
	_ = dc.CreateDocument("one.md", "")
	_ = dc.CreateDocument("two.md", "")
	docs, err := tb.Documents()
	if err != nil {
		t.Fatal(err)
	}
	actual, err := docs.FileNames()
	if err != nil {
		t.Fatal(err)
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
	tb, dc := test.NewDocumentCreator(t)
	doc1 := dc.CreateDocument("one.md", "# The one\n[The other](two.md)")
	doc2 := dc.CreateDocument("two.md", "# The other\n[The one](one.md)")
	docs, err := tb.Documents()
	if err != nil {
		t.Fatal(err)
	}
	actual, err := docs.TikiLinks()
	if err != nil {
		t.Fatal(err)
	}
	expected := domain.TikiLinkCollection{
		domain.NewTikiLink("The other", doc1.TitleSection(), doc2),
		domain.NewTikiLink("The one", doc2.TitleSection(), doc1),
	}
	diff := cmp.Diff(expected, actual, cmp.AllowUnexported(expected[0], expected[0].SourceSection(), expected[0].TargetDocument()))
	if diff != "" {
		t.Fatal(diff)
	}
}
