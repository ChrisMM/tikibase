package domain_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kevgo/tikibase/domain"
)

func TestTikiSectionAnchor(t *testing.T) {
	section := domain.ScaffoldTikiSection("### what is it\n")
	actual := section.Anchor()
	expected := "what-is-it"
	if actual != expected {
		t.Fatalf("mismatching section anchors: expected '%s', got '%s'", expected, actual)
	}
}

func TestTikiSectionContent(t *testing.T) {
	expectedContent := "### title\nthe content\n"
	section := domain.ScaffoldTikiSection(expectedContent)
	actualContent := string(section.Content())
	if actualContent != expectedContent {
		t.Fatalf("mismatching content! expected '%s' got '%s'", expectedContent, actualContent)
	}
}

func TestTikiSectionTikiLinks(t *testing.T) {
	docs := domain.ScaffoldDocumentCollection([]domain.DocumentScaffold{
		{FileName: "one.md", Content: "# One"},
		{FileName: "two.md", Content: `# Title\ntext [MD link to doc1](one.md)\n text [MD link to doc2](two.md) text\ntext <a href="one.md">HTML link to doc1</a> text <a textrun="dope">not a link</a>`},
	})
	section := docs[1].TitleSection()
	actual, err := section.TikiLinks(docs)
	if err != nil {
		t.Fatalf("cannot get links in section: %v", err)
	}
	expected := domain.ScaffoldTikiLinkCollection([]domain.TikiLinkScaffold{
		{Title: "MD link to doc1", SourceSection: section, TargetDocument: docs[0]},
		{Title: "MD link to doc2", SourceSection: section, TargetDocument: docs[1]},
		{Title: "HTML link to doc1", SourceSection: section, TargetDocument: docs[0]},
	})
	diff := cmp.Diff(expected, actual, cmp.AllowUnexported(expected[0], section, docs[0]))
	if diff != "" {
		t.Fatal(diff)
	}
}

func TestTikiSectionTitle(t *testing.T) {
	section := domain.ScaffoldTikiSection("### What is it\n")
	actual := section.Title()
	expected := "What is it"
	if actual != expected {
		t.Fatalf("mismatching section title: expected '%s', got '%s'", expected, actual)
	}
}
