package domain_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kevgo/tikibase/domain"
)

func TestSectionAnchor(t *testing.T) {
	section := domain.ScaffoldSection(domain.SectionScaffold{Content: "### what is it\n"})
	actual, err := section.Anchor()
	if err != nil {
		t.Fatal(err)
	}
	expected := "what-is-it"
	if actual != expected {
		t.Fatalf("mismatching section anchors: expected '%s', got '%s'", expected, actual)
	}
}

func TestSectionAppendLine(t *testing.T) {
	section := domain.ScaffoldSection(domain.SectionScaffold{Content: "existing content\n"})
	newSection := section.AppendLine("new line\n")
	actual := newSection.Content()
	expected := domain.SectionContent("existing content\nnew line\n")
	if actual != expected {
		t.Fatalf("did not append line correctly: expected '%s', got '%s'", expected, actual)
	}
}

func TestSectionContent(t *testing.T) {
	expectedContent := "### title\nthe content\n"
	section := domain.ScaffoldSection(domain.SectionScaffold{Content: expectedContent})
	actualContent := string(section.Content())
	if actualContent != expectedContent {
		t.Fatalf("mismatching content! expected '%s' got '%s'", expectedContent, actualContent)
	}
}

func TestSectionTikiLinks(t *testing.T) {
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
		{Title: "MD link to doc1", SourceSection: section, TargetDocument: &docs[0]},
		{Title: "MD link to doc2", SourceSection: section, TargetDocument: &docs[1]},
		{Title: "HTML link to doc1", SourceSection: section, TargetDocument: &docs[0]},
	})
	diff := cmp.Diff(expected, actual)
	if diff != "" {
		t.Fatal(diff)
	}
}

func TestSectionTikiLinksIgnoresHtmlLinks(t *testing.T) {
	docs := domain.ScaffoldDocumentCollection([]domain.DocumentScaffold{
		{FileName: "one.md", Content: "# Title\ntext [HTML link](http://google.com)"},
	})
	section := docs[0].TitleSection()
	actual, err := section.TikiLinks(docs)
	if err != nil {
		t.Fatalf("cannot get links in section: %v", err)
	}
	if len(actual) != 0 {
		t.Fatalf("shouldn't have found any links here")
	}
}

func TestSectionTitle(t *testing.T) {
	section := domain.ScaffoldSection(domain.SectionScaffold{Content: "### What is it\n"})
	actual, err := section.Title()
	if err != nil {
		t.Fatal(err)
	}
	expected := "What is it"
	if actual != expected {
		t.Fatalf("mismatching section title: expected '%s', got '%s'", expected, actual)
	}
}

func TestSectionURL(t *testing.T) {
	doc := domain.ScaffoldDocument(domain.DocumentScaffold{FileName: "one.md"})
	section := domain.ScaffoldSection(domain.SectionScaffold{Content: "### What is it\n", Doc: &doc})
	actual, err := section.URL()
	if err != nil {
		t.Fatal(err)
	}
	expected := "one.md#what-is-it"
	if actual != expected {
		t.Fatalf("mismatching section URL: expected '%s', got '%s'", expected, actual)
	}
}
