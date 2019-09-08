package domain_test

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/kevgo/tikibase/domain"
)

// NewTempDirectoryTikiBase provides an empty DirectoryTikiBase instance
// in the system's temp directory.
// Repeated calls to this return unique instances.
func newTempDirectoryTikiBase(t *testing.T) domain.DirectoryTikiBase {
	tmpDir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatal("cannot create temp dir")
	}
	result, err := domain.NewDirectoryTikiBase(tmpDir)
	if err != nil {
		t.Fatalf("cannot create new TikiBase: %v", err)
	}
	return result
}

func TestDirectoryTikiBaseCreateDocument(t *testing.T) {
	tb := newTempDirectoryTikiBase(t)
	_, err := tb.CreateDocument("one.md", "The one.")
	if err != nil {
		t.Fatalf("cannot create document: %v", err)
	}
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

func TestDirectoryTikiBaseDocuments(t *testing.T) {
	tb1 := newTempDirectoryTikiBase(t)
	_, err := tb1.CreateDocument("one", "")
	if err != nil {
		t.Fatalf("cannot create document 1")
	}
	_, err = tb1.CreateDocument("two", "")
	if err != nil {
		t.Fatalf("cannot create document 2")
	}

	// get the documents
	tb2, err := domain.NewDirectoryTikiBase(tb1.StorageDir())
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

func TestNewDirectoryTikiBase(t *testing.T) {
	_, err := domain.NewDirectoryTikiBase(".")
	if err != nil {
		t.Fatal(err)
	}
}
