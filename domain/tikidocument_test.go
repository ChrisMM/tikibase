package domain_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kevgo/tikibase/domain"
	"github.com/kevgo/tikibase/test"
)

// newTempTikiDocument provides a TikiDocument instance inside a TikiBase in a temp directory.
func newTempTikiDocument(filename domain.TikiDocumentFilename, content string, t *testing.T) (domain.TikiBase, domain.TikiDocument) {
	tb := test.NewTempTikiBase(t)
	td, err := tb.CreateDocument(filename, content)
	if err != nil {
		t.Fatal(err)
	}
	return tb, td
}

func TestTikiDocumentAllSections(t *testing.T) {
	_, td := newTempTikiDocument("one.md", "# Title\nmy doc\n### One\nThe one.\n### Two\nThe other.", t)
	// TODO: compare against expected datastructure
	sections := td.AllSections()
	if len(sections) != 3 {
		t.Fatalf("unexpected sections length: expected 3 got %d", len(sections))
	}

	// verify title section
	expected := domain.TikiSectionContent("# Title\nmy doc")
	actual := sections[0].Content()
	if actual != expected {
		t.Fatalf("unexpected title section: expected '%s' got '%s'", expected, actual)
	}

	// verify content section 1
	expected = "### One\nThe one."
	actual = sections[1].Content()
	if actual != expected {
		t.Fatalf("unexpected content section 1: expected '%s' got '%s'", expected, actual)
	}

	// verify content section 2
	expected = "### Two\nThe other."
	actual = sections[2].Content()
	if actual != expected {
		t.Fatalf("unexpected content section 2: expected '%s' got '%s'", expected, actual)
	}
}

func TestTikiDocumentFileName(t *testing.T) {
	tb := test.NewTempTikiBase(t)
	td, err := tb.CreateDocument("one.md", "")
	if err != nil {
		t.Fatalf("cannot create document one.md: %v", err)
	}
	expectedFileName := "one.md"
	actualFileName := string(td.FileName())
	if actualFileName != expectedFileName {
		t.Fatalf("expected '%s' got '%s'", expectedFileName, actualFileName)
	}
}

func TestTikiDocumentTikiLinks(t *testing.T) {
	docs := domain.ScaffoldTikiDocumentCollection([]domain.TikiDocumentScaffold{
		{FileName: "doc1.md", Content: "### One\n"},
		{FileName: "doc2.md", Content: "### Two\n[one](doc1.md)"},
	})
	actual, err := docs[1].TikiLinks(docs)
	if err != nil {
		t.Fatalf("error getting TikiLinks for doc2: %v", err)
	}
	expected := domain.ScaffoldTikiLinkCollection([]domain.TikiLinkScaffold{
		{Title: "one", SourceSection: docs[1].TitleSection(), TargetDocument: docs[0]},
	})
	diff := cmp.Diff(expected, actual, cmp.AllowUnexported(expected[0], docs[0], docs[1].TitleSection()))
	if diff != "" {
		t.Fatal(diff)
	}
}

func TestTikiDocumentTitleSection(t *testing.T) {
	_, doc := newTempTikiDocument(domain.TikiDocumentFilename("one.md"), "# My Title\n\nTitle section content.\n\n### Content Section\n Content section content.\n", t)
	section := doc.TitleSection()
	expectedContent := "# My Title\n\nTitle section content.\n"
	diff := cmp.Diff(string(section.Content()), expectedContent)
	if diff != "" {
		t.Fatalf("mismatching section content: %s", diff)
	}
}
