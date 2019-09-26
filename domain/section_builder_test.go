package domain_test

import (
	"testing"

	"github.com/kevgo/tikibase/domain"
)

func TestSectionBuilder(t *testing.T) {
	doc := domain.ScaffoldDocument(domain.DocumentScaffold{})
	tsb := domain.NewSectionBuilder("# Title\n", doc)
	tsb.AddLine("\n")
	tsb.AddLine("content 1\n")
	tsb.AddLine("content 2\n")
	tsb.AddLine("\n")
	section := tsb.Section()
	expectedContent := domain.SectionContent("# Title\n\ncontent 1\ncontent 2\n\n")
	actualContent := section.Content()
	if actualContent != expectedContent {
		t.Fatalf("TikiSectionBuilder didn't build the right content!\nEXPECTED:\n'%s'\n ACTUAL:\n'%s'", expectedContent, actualContent)
	}
	if section.Document() != doc {
		t.Fatalf("Created section doesn't contain a link to its containing document")
	}
}
