package storage_test

import (
	"testing"

	"github.com/kevgo/tikibase/storage"
)

func TestTikiSectionContent(t *testing.T) {
	expectedContent := "the content"
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
	section := storage.NewTikiSection("# Title [title 1](one.md) text [title 2](two.md)")
	links, err := section.TikiLinks(documents)
	if err != nil {
		t.Fatalf("cannot get links in section: %v", err)
	}
	if len(links) != 2 {
		t.Fatalf("unexpected link count: expected 2, got %d", len(links))
	}
	if links[0].TargetDocument() != doc1 {
		t.Fatal("unexpected target document")
	}
	if links[0].SourceSection() != section {
		t.Fatal("unexpected source section")
	}
	if links[1].TargetDocument() != doc2 {
		t.Fatal("unexpected target document")
	}
	if links[1].SourceSection() != section {
		t.Fatal("unexpected source section")
	}
}
