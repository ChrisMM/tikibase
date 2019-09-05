package storage_test

import (
	"io/ioutil"
	"testing"

	"github.com/kevgo/tikibase/storage"
)

func TestNewTikiBase(t *testing.T) {
	_, err := storage.NewTikiBase("foo")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDocuments(t *testing.T) {
	tb1, err := createTestBase(t)
	if err != nil {
		t.Fatal(err)
	}
	_, err = tb1.CreateDocument("one", "")
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

func TestSaveDocument(t *testing.T) {

}

// createTestBase creates a test TikiBase in a temp directory
func createTestBase(t *testing.T) (storage.TikiBase, error) {
	tmpDir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatal("cannot create temp dir")
	}
	return storage.NewTikiBase(tmpDir)
}
