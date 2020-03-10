package occurrences_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/kevgo/tikibase/domain"
	"github.com/kevgo/tikibase/occurrences"
)

func TestRenderOccurrencesSection(t *testing.T) {
	targetDoc := domain.ScaffoldDocument(domain.DocumentScaffold{FileName: "programming-language.md", Content: "# Programming Language\n### what is it\n- system to author software\n"})
	goDoc := domain.ScaffoldDocument(domain.DocumentScaffold{FileName: "go.md", Content: "# Go\n### what is it\n- a [programming language](programming-language.md)\n"})
	tsDoc := domain.ScaffoldDocument(domain.DocumentScaffold{FileName: "typescript.md", Content: "# TypeScript\n### what is it\n- a [programming language](programming-language.md)\n"})
	linksToDoc := domain.ScaffoldTikiLinkCollection([]domain.TikiLinkScaffold{
		{Title: "programming language", SourceSection: goDoc.ContentSections()[0], TargetDocument: targetDoc},
		{Title: "programming language", SourceSection: tsDoc.ContentSections()[0], TargetDocument: targetDoc},
	})

	renderedSection, err := occurrences.CreateOccurrencesSection(linksToDoc, targetDoc)

	if err != nil {
		t.Fatal(err)
	}
	expected := `### occurrences

- [Go (what is it)](go.md#what-is-it)
- [TypeScript (what is it)](typescript.md#what-is-it)
`
	actual := string(renderedSection.Content())
	diff := cmp.Diff(expected, actual)
	if diff != "" {
		t.Fatal(diff)
	}
}

func TestRenderOccurrencesSection_LinkToTitleSection(t *testing.T) {
	targetDoc := domain.ScaffoldDocument(domain.DocumentScaffold{FileName: "programming-language.md", Content: "# Programming Language\n### what is it\n- system to author software\n"})
	goDoc := domain.ScaffoldDocument(domain.DocumentScaffold{FileName: "go.md", Content: "# Go\na [programming language](programming-language.md)\n"})
	linksToDoc := domain.ScaffoldTikiLinkCollection([]domain.TikiLinkScaffold{
		{Title: "programming language", SourceSection: &(*goDoc.TitleSection()), TargetDocument: targetDoc},
	})

	renderedSection, err := occurrences.CreateOccurrencesSection(linksToDoc, targetDoc)

	if err != nil {
		t.Fatal(err)
	}
	expected := `### occurrences

- [Go](go.md)
`
	actual := string(renderedSection.Content())
	diff := cmp.Diff(expected, actual)
	if diff != "" {
		t.Fatal(diff)
	}
}
