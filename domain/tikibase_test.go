package domain_test

import (
	"os"
	"path"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kevgo/tikibase/domain"
	"github.com/kevgo/tikibase/test"
)

func TestTikiBaseCreateDocument(t *testing.T) {
	tb := test.NewTempTikiBase(t)
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

func TestTikiBaseDocuments(t *testing.T) {
	tb1, dc := test.NewDocumentCreator(t)
	_ = dc.CreateDocument("one", "")
	_ = dc.CreateDocument("two", "")

	// get the documents
	tb2, err := domain.NewTikiBase(tb1.StorageDir())
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

func TestTikiBaseLoad(t *testing.T) {
	tb := test.NewTempTikiBase(t)
	expected, err := tb.CreateDocument("one.md", "The one")
	if err != nil {
		t.Fatal(err)
	}
	tb2, err := domain.NewTikiBase(tb.StorageDir())
	if err != nil {
		t.Fatal(err)
	}
	actual, err := tb2.Load("one.md")
	if err != nil {
		t.Fatal(err)
	}
	diff := cmp.Diff(expected, actual, cmp.AllowUnexported(actual))
	if diff != "" {
		t.Fatal(diff)
	}
}

func TestNewTikiBase(t *testing.T) {
	_, err := domain.NewTikiBase(".")
	if err != nil {
		t.Fatal(err)
	}
}
