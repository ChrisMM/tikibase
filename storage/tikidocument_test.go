package storage_test

import (
	"testing"

	"github.com/kevgo/tikibase/storage"
)

func TestAllSections(t *testing.T) {
	td := storage.NewTikiDocument("handle", "# Title\nmy doc\n### One\nThe one.\n### Two\nThe other.")
	sections := td.AllSections()
	if len(sections) != 3 {
		t.Fatalf("unexpected sections length: expected 3 got %d", len(sections))
	}

	// verify title section
	expected := storage.TikiSectionContent("# Title\nmy doc\n")
	actual := sections[0].Content()
	if actual != expected {
		t.Fatalf("unexpected title section: expected '%s' got '%s'", expected, actual)
	}

	// verify content section 1
	expected = "### One\nThe one.\n"
	actual = sections[1].Content()
	if actual != expected {
		t.Fatalf("unexpected content section 1: expected '%s' got '%s'", expected, actual)
	}

	// verify content section 2
	expected = "### Two\nThe other.\n"
	actual = sections[2].Content()
	if actual != expected {
		t.Fatalf("unexpected content section 2: expected '%s' got '%s'", expected, actual)
	}
}

func TestTikiDocumentFilePath(t *testing.T) {
	td := storage.NewTikiDocument(storage.Handle("one"), "")
	expectedFilePath := "one.md"
	actualFilePath := td.FilePath()
	if actualFilePath != expectedFilePath {
		t.Fatalf("expected '%s' got '%s'", expectedFilePath, actualFilePath)
	}
}

func TestHandle(t *testing.T) {
	expectedHandle := storage.Handle("handle")
	td := storage.NewTikiDocument(expectedHandle, "content")
	actualHandle := td.Handle()
	if actualHandle != expectedHandle {
		t.Fatalf("mismatching handle. expected '%s' got '%s'", expectedHandle, actualHandle)
	}
}
