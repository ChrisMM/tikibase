package domain_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/kevgo/tikibase/domain"
)

func TestSectionCollectionEqual(t *testing.T) {
	doc1 := domain.ScaffoldDocument(domain.DocumentScaffold{FileName: "one.md"})
	doc2 := domain.ScaffoldDocument(domain.DocumentScaffold{FileName: "two.md"})
	expected := domain.ScaffoldSectionCollection([]domain.SectionScaffold{
		{Content: "section 1\n", Doc: &doc1},
		{Content: "section 2\n", Doc: &doc2},
	})

	// compare against SectionCollection with similar content
	match := domain.ScaffoldSectionCollection([]domain.SectionScaffold{
		{Content: "section 1\n", Doc: &doc1},
		{Content: "section 2\n", Doc: &doc2},
	})
	if diff := cmp.Diff(expected, match); diff != "" {
		t.Fatalf("match: unexpected mismatch: %s", diff)
	}

	// compare against shorter SectionCollection
	shorter := domain.ScaffoldSectionCollection([]domain.SectionScaffold{
		{Content: "section 1\n", Doc: &doc1},
	})
	if diff := cmp.Diff(expected, shorter); diff == "" {
		t.Fatalf("shorter: unexpected match: %s", diff)
	}

	// compare against longer SectionCollection
	longer := domain.ScaffoldSectionCollection([]domain.SectionScaffold{
		{Content: "section 1\n", Doc: &doc1},
		{Content: "section 2\n", Doc: &doc2},
		{Content: "section 1\n", Doc: &doc1},
	})
	if diff := cmp.Diff(expected, longer); diff == "" {
		t.Fatalf("longer: unexpected match: %s", diff)
	}

	// compare against SectionCollection with different text content
	differentContent := domain.ScaffoldSectionCollection([]domain.SectionScaffold{
		{Content: "section A\n", Doc: &doc1},
		{Content: "section 2\n", Doc: &doc2},
	})
	if diff := cmp.Diff(expected, differentContent); diff == "" {
		t.Fatalf("differentContent: unexpected match: %s", diff)
	}

	// compare against SectionCollection with different document
	differentDoc := domain.ScaffoldSectionCollection([]domain.SectionScaffold{
		{Content: "section 1\n", Doc: &doc1},
		{Content: "section 2\n", Doc: &doc1},
	})
	if diff := cmp.Diff(expected, differentDoc); diff == "" {
		t.Fatalf("differentDoc: unexpected match: %s", diff)
	}
}

func TestSectionCollectionReplace(t *testing.T) {
	sections := domain.ScaffoldSectionCollection([]domain.SectionScaffold{
		{Content: "section 1\n"},
		{Content: "section 2\n"},
	})
	newSection2 := domain.ScaffoldSection(domain.SectionScaffold{Content: "new section 2\n"})
	actual := sections.Replace(sections[1], newSection2)
	expected := domain.SectionCollection([]domain.Section{sections[0], newSection2})
	if diff := cmp.Diff(expected, actual); diff != "" {
		t.Fatal(diff)
	}
}

func TestSectionCollectionText(t *testing.T) {
	sections := domain.ScaffoldSectionCollection([]domain.SectionScaffold{
		{Content: "section 1\n"},
		{Content: "section 2\n"},
	})
	actual := sections.Text()
	expected := "section 1\nsection 2\n"
	if actual != expected {
		t.Fatalf("mismatching content: expected '%s', got '%s'", expected, actual)
	}
}
