package domain_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kevgo/tikibase/domain"
)

func TestDocumentAllSections(t *testing.T) {
	td := domain.ScaffoldDocument(domain.DocumentScaffold{
		FileName: "one.md", Content: "# Title\nmy doc\n### One\nThe one.\n### Two\nThe other.",
	})
	// TODO: compare against expected datastructure
	sections := td.AllSections()
	if len(sections) != 3 {
		t.Fatalf("unexpected sections length: expected 3 got %d", len(sections))
	}

	// verify title section
	expected := domain.SectionContent("# Title\nmy doc")
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

func TestDocumentContentSections(t *testing.T) {
	td := domain.ScaffoldDocument(domain.DocumentScaffold{
		FileName: "one.md", Content: "# Title\nmy doc\n### One\nThe one.\n### Two\nThe other.",
	})
	// TODO: compare against expected datastructure
	sections := td.ContentSections()
	if len(sections) != 2 {
		t.Fatalf("unexpected sections length: expected 2 got %d", len(sections))
	}

	// verify content section 1
	expected := "### One\nThe one."
	actual := string(sections[0].Content())
	if actual != expected {
		t.Fatalf("unexpected content section 1: expected '%s' got '%s'", expected, actual)
	}

	// verify content section 2
	expected = "### Two\nThe other."
	actual = string(sections[1].Content())
	if actual != expected {
		t.Fatalf("unexpected content section 2: expected '%s' got '%s'", expected, actual)
	}
}

func TestDocumentFileName(t *testing.T) {
	doc := domain.ScaffoldDocument(domain.DocumentScaffold{FileName: "one.md"})
	expectedFileName := "one.md"
	actualFileName := string(doc.FileName())
	if actualFileName != expectedFileName {
		t.Fatalf("expected '%s' got '%s'", expectedFileName, actualFileName)
	}
}

func TestDocumentTikiLinks(t *testing.T) {
	docs := domain.ScaffoldDocumentCollection([]domain.DocumentScaffold{
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

func TestDocumentTitleSection(t *testing.T) {
	doc := domain.ScaffoldDocument(domain.DocumentScaffold{Content: "# My Title\n\nTitle section content.\n\n### Content Section\n Content section content.\n"})
	section := doc.TitleSection()
	expectedContent := "# My Title\n\nTitle section content.\n"
	diff := cmp.Diff(string(section.Content()), expectedContent)
	if diff != "" {
		t.Fatalf("mismatching section content: %s", diff)
	}
}

func TestScaffoldDocument(t *testing.T) {
	actual := domain.ScaffoldDocument(domain.DocumentScaffold{})
	if actual.FileName() == "" {
		t.Fatal("no default FileName")
	}
	if actual.TitleSection().Title() == "" {
		t.Fatal("no default section")
	}
}
