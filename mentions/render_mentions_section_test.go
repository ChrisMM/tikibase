package mentions_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/kevgo/tikibase/domain"
	"github.com/kevgo/tikibase/mentions"
)

func TestRenderMentionsSection(t *testing.T) {
	targetDoc := domain.ScaffoldDocument(domain.DocumentScaffold{FileName: "programming-language.md", Content: "# Programming Language\n### what is it\n- system to author software\n"})
	goDoc := domain.ScaffoldDocument(domain.DocumentScaffold{FileName: "go.md", Content: "# Go\n### what is it\n- a [programming language](programming-language.md)\n"})
	tsDoc := domain.ScaffoldDocument(domain.DocumentScaffold{FileName: "typescript.md", Content: "# TypeScript\n### what is it\n- a [programming language](programming-language.md)\n"})
	linksToDoc := domain.ScaffoldTikiLinkCollection([]domain.TikiLinkScaffold{
		{Title: "programming language", SourceSection: goDoc.ContentSections()[0], TargetDocument: targetDoc},
		{Title: "programming language", SourceSection: tsDoc.ContentSections()[0], TargetDocument: targetDoc},
	})
	renderedSection := mentions.RenderMentionsSection(linksToDoc, &targetDoc)
	expected := `### mentions

- [Go](go.md#what-is-it)
- [TypeScript](typescript.md#what-is-it)
`
	actual := string(renderedSection.Content())
	diff := cmp.Diff(expected, actual)
	if diff != "" {
		t.Fatal(diff)
	}
}
