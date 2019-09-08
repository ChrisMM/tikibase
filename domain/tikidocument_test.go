package domain_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kevgo/tikibase/domain"
)

// newTempTikiDocument provides a TikiDocument instance inside a DirectoryTikiBase in a temp directory.
func newTempTikiDocument(filename domain.TikiDocumentFilename, content string, t *testing.T) (domain.TikiBase, domain.TikiDocument) {
	tb := newTempDirectoryTikiBase(t)
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
	tb := newTempDirectoryTikiBase(t)
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

func TestTikiDocumentLinks(t *testing.T) {
	tb := newTempDirectoryTikiBase(t)
	doc1, err := tb.CreateDocument("doc1.md", "### One\n")
	if err != nil {
		t.Fatalf("cannot created doc1: %v", err)
	}
	doc2, err := tb.CreateDocument("doc2.md", "### Two\n[one](doc1.md)")
	if err != nil {
		t.Fatalf("cannot created doc2: %v", err)
	}
	actual, err := doc2.TikiLinks(tb)
	if err != nil {
		t.Fatalf("cannot get TikiLinks for doc2: %v", err)
	}
	expected := []domain.TikiLink{
		domain.NewTikiLink("one", doc2.TitleSection(), doc1),
	}
	diff := cmp.Diff(expected, actual, cmp.AllowUnexported(expected[0], doc1, doc2.TitleSection()))
	if diff != "" {
		t.Fatal(diff)
	}
}
