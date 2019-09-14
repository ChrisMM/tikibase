package domain_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kevgo/tikibase/domain"
)

func TestDocumentAllSections(t *testing.T) {
	td := domain.ScaffoldDocument(domain.DocumentScaffold{
		FileName: "one.md", Content: "# Title\n\nmy doc\n\n### One\n\nThe one.\n\n### Two\n\nThe other.\n",
	})
	// TODO: compare against expected datastructure
	sections := td.AllSections()
	if len(sections) != 3 {
		t.Fatalf("unexpected sections length: expected 3 got %d", len(sections))
	}

	// verify title section
	expected := domain.SectionContent("# Title\n\nmy doc\n\n")
	actual := sections[0].Content()
	if actual != expected {
		t.Fatalf("unexpected title section: expected '%s' got '%s'", expected, actual)
	}

	// verify content section 1
	expected = "### One\n\nThe one.\n\n"
	actual = sections[1].Content()
	if actual != expected {
		t.Fatalf("unexpected content section 1: expected '%s' got '%s'", expected, actual)
	}

	// verify content section 2
	expected = "### Two\n\nThe other.\n"
	actual = sections[2].Content()
	if actual != expected {
		t.Fatalf("unexpected content section 2: expected '%s' got '%s'", expected, actual)
	}
}

func TestDocumentAppendSection(t *testing.T) {
	oldDoc := domain.ScaffoldDocument(domain.DocumentScaffold{FileName: "one.md", Content: "existing document content\n"})
	newSection := domain.ScaffoldSection(domain.SectionScaffold{Content: "### new section\n"})
	newDoc := oldDoc.AppendSection(newSection)
	expectedContent := "existing document content\n\n### new section\n"
	if newDoc.Content() != expectedContent {
		t.Fatalf("mismatching document content: expected '%s', got '%s'", expectedContent, newDoc.Content())
	}
	expectedFileName := domain.DocumentFilename("one.md")
	if newDoc.FileName() != expectedFileName {
		t.Fatalf("didn't bring the filename over to the new doc: expected '%s', got '%s'", expectedFileName, newDoc.FileName())
	}
}

func TestDocumentContentSections(t *testing.T) {
	td := domain.ScaffoldDocument(domain.DocumentScaffold{
		FileName: "one.md", Content: "# Title\nmy doc\n### One\nThe one.\n### Two\nThe other.\n",
	})
	// TODO: compare against expected datastructure
	sections := td.ContentSections()
	if len(sections) != 2 {
		t.Fatalf("unexpected sections length: expected 2 got %d", len(sections))
	}

	// verify content section 1
	expected := domain.SectionContent("### One\nThe one.\n")
	actual := sections[0].Content()
	if actual != expected {
		t.Fatalf("unexpected content section 1: expected '%s' got '%s'", expected, actual)
	}

	// verify content section 2
	expected = domain.SectionContent("### Two\nThe other.\n")
	actual = sections[1].Content()
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

func TestDocumentReplaceSection(t *testing.T) {
	td := domain.ScaffoldDocument(domain.DocumentScaffold{
		FileName: "one.md", Content: "# Title\n\nmy doc\n\n### One\n\nThe one.\n\n### Two\n\nThe other.\n",
	})
	sections := td.AllSections()
	twoSection := sections[2]
	newSection := domain.ScaffoldSection(domain.SectionScaffold{Content: "### Two\n\nThe second.\n", Doc: &td})
	newdoc := td.ReplaceSection(twoSection, newSection)

	newSections := newdoc.AllSections()
	if len(newSections) != 3 {
		t.Fatalf("unexpected newSections length: expected 3 got %d", len(newSections))
	}

	// verify title section
	expected := domain.SectionContent("# Title\n\nmy doc\n\n")
	actual := newSections[0].Content()
	if actual != expected {
		t.Fatalf("unexpected title section: expected '%s' got '%s'", expected, actual)
	}

	// verify content section 1
	expected = domain.SectionContent("### One\n\nThe one.\n\n")
	actual = newSections[1].Content()
	if actual != expected {
		t.Fatalf("unexpected content section 1: expected '%s' got '%s'", expected, actual)
	}

	// verify content section 2
	expected = domain.SectionContent("### Two\n\nThe second.\n")
	actual = newSections[2].Content()
	if actual != expected {
		t.Fatalf("unexpected content section 2: expected '%s' got '%s'", expected, actual)
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
	expectedContent := "# My Title\n\nTitle section content.\n\n"
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
