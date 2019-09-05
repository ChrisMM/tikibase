package storage_test

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/kevgo/tikibase/storage"
)

func TestCreateDocument(t *testing.T) {
	tb := createTestBase(t)
	tb.CreateDocument("one", "The one.")
	filePath := path.Join(tb.StorageDir(), "one.md")
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		t.Fatalf("file %s not found: %v", filePath, err)
	}
	if fileInfo.IsDir() {
		t.Fatalf("file %s should not be a directory", filePath)
	}
	if fileInfo.Mode() != 0644 {
		t.Fatalf("file %s should have access 0644 but has %#o", filePath, fileInfo.Mode())
	}
}

func TestDocuments(t *testing.T) {
	tb1 := createTestBase(t)
	_, err := tb1.CreateDocument("one", "")
	if err != nil {
		t.Fatalf("cannot create document 1")
	}
	_, err = tb1.CreateDocument("two", "")
	if err != nil {
		t.Fatalf("cannot create document 2")
	}

	// get the documents
	tb2, err := storage.NewTikiBase(tb1.StorageDir())
	if err != nil {
		t.Fatalf("cannot instantiate tb2: %v", err)
	}
	actual, err := tb2.Documents()
	if err != nil {
		t.Fatalf("cannot call tb.Documents(): %v", err)
	}

	// verify results
	if len(actual) != 2 {
		t.Errorf("expected %d documents but got %d", 2, len(actual))
	}
}

func TestNewTikiBase(t *testing.T) {
	_, err := storage.NewTikiBase(".")
	if err != nil {
		t.Fatal(err)
	}
}

func TestSaveDocument(t *testing.T) {
	tb := createTestBase(t)
	expectedContent := "The content."
	doc := storage.TikiDocument{Handle: "my-handle", Content: expectedContent}
	err := tb.SaveDocument(doc)
	if err != nil {
		t.Fatalf("cannot save document: %v", err)
	}
	filePath := path.Join(tb.StorageDir(), "my-handle.md")
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		t.Fatalf("file %s not found: %v", filePath, err)
	}
	if fileInfo.IsDir() {
		t.Fatalf("file %s should not be a directory", filePath)
	}
	if fileInfo.Mode() != 0644 {
		t.Fatalf("file %s should have access 0644 but has %#o", filePath, fileInfo.Mode())
	}
	actualContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Fatalf("cannot read file %s: %v", filePath, err)
	}
	if string(actualContent) != expectedContent {
		t.Fatalf("diverging file content. Expected '%s', found '%s'", expectedContent, string(actualContent))
	}
}

// createTestBase creates a test TikiBase in a temp directory
func createTestBase(t *testing.T) storage.TikiBase {
	tmpDir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatal("cannot create temp dir")
	}
	result, err := storage.NewTikiBase(tmpDir)
	if err != nil {
		t.Fatalf("cannot create new TikiBase: %v", err)
	}
	return result
}
