package domain_test

import (
	"testing"

	"github.com/kevgo/tikibase/domain"
)

func TestSectionBuilder(t *testing.T) {
	doc := domain.ScaffoldDocument(domain.DocumentScaffold{})
	tsb := domain.NewSectionBuilder("# Title", &doc)
	tsb.AddLine("content 1")
	tsb.AddLine("content 2")
	tsb.AddLine("")
	section := tsb.Section()
	expectedContent := domain.SectionContent("# Title\ncontent 1\ncontent 2\n")
	actualContent := section.Content()
	if actualContent != expectedContent {
		t.Fatalf("TikiSectionBuilder didn't build the right content! expected '%s' got '%s'", expectedContent, actualContent)
	}
	if section.Document() != &doc {
		t.Fatalf("Created section doesn't contain a link to its containing document")
	}
}
