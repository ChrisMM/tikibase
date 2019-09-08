package domain_test

import (
	"testing"

	"github.com/kevgo/tikibase/domain"
)

func TestTikiSectionBuilder(t *testing.T) {
	tsb := domain.NewTikiSectionBuilder("# Title")
	tsb.AddLine("content 1")
	tsb.AddLine("content 2")
	section := tsb.Section()
	expectedContent := domain.TikiSectionContent("# Title\ncontent 1\ncontent 2\n")
	actualContent := section.Content()
	if actualContent != expectedContent {
		t.Fatalf("TikiSectionBuilder didn't build the right content! expected '%s' got '%s'", expectedContent, actualContent)
	}
}
