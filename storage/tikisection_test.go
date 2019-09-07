package storage_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kevgo/tikibase/storage"
)

func TestTikiSectionContent(t *testing.T) {
	expectedContent := storage.TikiSectionContent("### title\nthe content")
	ts := storage.NewTikiSection(expectedContent)
	actualContent := ts.Content()
	if ts.Content() != expectedContent {
		t.Fatalf("mismatching content! expected '%s' got '%s'", expectedContent, actualContent)
	}
}

func TestTikiSectionLinks(t *testing.T) {
	documents := storage.NewTikiDocumentCollection()
	doc1 := storage.NewTikiDocument("one", "# One")
	doc2 := storage.NewTikiDocument("two", "# Two")
	documents.Add(doc1, doc2)
	section := storage.NewTikiSection(`# Title\ntext [MD link to doc1](one.md)\n text [MD link to doc2](two.md) text\ntext <a href="one.md">HTML link to doc1</a> text <a textrun="dope">not a link</a>`)
	links, err := section.TikiLinks(documents)
	if err != nil {
		t.Fatalf("cannot get links in section: %v", err)
	}
	expected := []storage.TikiLink{
		storage.NewTikiLink("MD link to doc1", section, doc1),
		storage.NewTikiLink("MD link to doc2", section, doc2),
		storage.NewTikiLink("HTML link to doc1", section, doc1),
	}
	diff := cmp.Diff(expected, links, cmp.AllowUnexported(expected[0], section, doc1))
	if diff != "" {
		t.Fatal(diff)
	}
}
