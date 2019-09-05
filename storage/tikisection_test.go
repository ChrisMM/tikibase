package storage_test

import (
	"testing"

	"github.com/kevgo/tikibase/storage"
)

func TestTikiSectionContent(t *testing.T) {
	expectedContent := "the content"
	ts := storage.NewTikiSection(expectedContent)
	actualContent := ts.Content()
	if ts.Content() != expectedContent {
		t.Fatalf("mismatching content! expected '%s' got '%s'", expectedContent, actualContent)
	}
}
