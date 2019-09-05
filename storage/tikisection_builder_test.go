package storage_test

import (
	"testing"

	"github.com/kevgo/tikibase/storage"
)

func TestTikiSectionBuilder(t *testing.T) {
	tsb := storage.NewTikiSectionBuilder("# Title")
	tsb.AddLine("content 1")
	tsb.AddLine("content 2")
	section := tsb.Section()
	expectedContent := "# Title\ncontent 1\ncontent 2\n"
	actualContent := section.Content()
	if actualContent != expectedContent {
		t.Fatalf("TikiSectionBuilder didn't build the right content! expected '%s' got '%s'", expectedContent, actualContent)
	}
}
