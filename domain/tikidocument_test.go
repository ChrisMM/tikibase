package domain_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kevgo/tikibase/domain"
)

func TestTikiDocumentAllSections(t *testing.T) {
	td := domain.NewTikiDocument("handle", "# Title\nmy doc\n### One\nThe one.\n### Two\nThe other.")
	sections := td.AllSections()
	if len(sections) != 3 {
		t.Fatalf("unexpected sections length: expected 3 got %d", len(sections))
	}

	// verify title section
	expected := domain.TikiSectionContent("# Title\nmy doc\n")
	actual := sections[0].Content()
	if actual != expected {
		t.Fatalf("unexpected title section: expected '%s' got '%s'", expected, actual)
	}

	// verify content section 1
	expected = "### One\nThe one.\n"
	actual = sections[1].Content()
	if actual != expected {
		t.Fatalf("unexpected content section 1: expected '%s' got '%s'", expected, actual)
	}

	// verify content section 2
	expected = "### Two\nThe other.\n"
	actual = sections[2].Content()
	if actual != expected {
		t.Fatalf("unexpected content section 2: expected '%s' got '%s'", expected, actual)
	}
}

func TestTikiDocumentFilePath(t *testing.T) {
	td := domain.NewTikiDocument(domain.Handle("one"), "")
	expectedFilePath := "one.md"
	actualFilePath := td.FilePath()
	if actualFilePath != expectedFilePath {
		t.Fatalf("expected '%s' got '%s'", expectedFilePath, actualFilePath)
	}
}

func TestTikiDocumentHandle(t *testing.T) {
	expectedHandle := domain.Handle("handle")
	td := domain.NewTikiDocument(expectedHandle, "content")
	actualHandle := td.Handle()
	if actualHandle != expectedHandle {
		t.Fatalf("mismatching handle. expected '%s' got '%s'", expectedHandle, actualHandle)
	}
}

func TestTikiDocumentLinks(t *testing.T) {
	docs := domain.NewTikiDocumentCollection()
	doc1 := domain.NewTikiDocument("doc1", "### One\n")
	doc2 := domain.NewTikiDocument("doc2", "### Two\n[one](doc1.md)")
	docs.Add(doc1, doc2)
	actual, err := doc2.TikiLinks(docs)
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
