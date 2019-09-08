package domain_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kevgo/tikibase/domain"
)

func TestTikiSectionContent(t *testing.T) {
	expectedContent := domain.TikiSectionContent("### title\nthe content")
	ts := domain.NewTikiSection(expectedContent)
	actualContent := ts.Content()
	if ts.Content() != expectedContent {
		t.Fatalf("mismatching content! expected '%s' got '%s'", expectedContent, actualContent)
	}
}

func TestTikiSectionLinks(t *testing.T) {
	documents := domain.NewTikiDocumentCollection()
	doc1 := domain.NewTikiDocument("one", "# One")
	doc2 := domain.NewTikiDocument("two", "# Two")
	documents.Add(doc1, doc2)
	section := domain.NewTikiSection(`# Title\ntext [MD link to doc1](one.md)\n text [MD link to doc2](two.md) text\ntext <a href="one.md">HTML link to doc1</a> text <a textrun="dope">not a link</a>`)
	links, err := section.TikiLinks(documents)
	if err != nil {
		t.Fatalf("cannot get links in section: %v", err)
	}
	expected := []domain.TikiLink{
		domain.NewTikiLink("MD link to doc1", section, doc1),
		domain.NewTikiLink("MD link to doc2", section, doc2),
		domain.NewTikiLink("HTML link to doc1", section, doc1),
	}
	diff := cmp.Diff(expected, links, cmp.AllowUnexported(expected[0], section, doc1))
	if diff != "" {
		t.Fatal(diff)
	}
}
