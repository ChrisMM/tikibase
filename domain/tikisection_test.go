package domain_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kevgo/tikibase/domain"
)

func newTestSection(content string, t *testing.T) domain.TikiSection {
	_, doc := newTempTikiDocument("one.md", content, t)
	return doc.TitleSection()
}

func TestTikiSectionAnchor(t *testing.T) {
	section := newTestSection("### what is it\n", t)
	actual := section.Anchor()
	expected := "what-is-it"
	if actual != expected {
		t.Fatalf("mismatching section anchors: expected '%s', got '%s'", expected, actual)
	}
}

func TestTikiSectionContent(t *testing.T) {
	expectedContent := "### title\nthe content\n"
	section := newTestSection(expectedContent, t)
	actualContent := string(section.Content())
	if actualContent != expectedContent {
		t.Fatalf("mismatching content! expected '%s' got '%s'", expectedContent, actualContent)
	}
}

func TestTikiSectionTikiLinks(t *testing.T) {
	tb := newTempDirectoryTikiBase(t)
	doc1, err := tb.CreateDocument("one.md", "# One")
	if err != nil {
		t.Fatalf("cannot create one.md: %v", err)
	}
	doc2, err := tb.CreateDocument("two.md", `# Title\ntext [MD link to doc1](one.md)\n text [MD link to doc2](two.md) text\ntext <a href="one.md">HTML link to doc1</a> text <a textrun="dope">not a link</a>`)
	if err != nil {
		t.Fatalf("cannot create two.md: %v", err)
	}
	section := doc2.TitleSection()
	links, err := section.TikiLinks(tb)
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

func TestTikiSectionTitle(t *testing.T) {
	section := newTestSection("### What is it\n", t)
	actual := section.Title()
	expected := "What is it"
	if actual != expected {
		t.Fatalf("mismatching section title: expected '%s', got '%s'", expected, actual)
	}
}
