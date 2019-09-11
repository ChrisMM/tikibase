package domain_test

import (
	"os"
	"path"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kevgo/tikibase/domain"
	"github.com/kevgo/tikibase/test"
)

func TestDirectoryTikiBaseCreateDocument(t *testing.T) {
	tb := test.NewTempDirectoryTikiBase(t)
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

func TestDirectoryTikiBaseDocumentFileNames(t *testing.T) {
	tb, dc := test.NewDocumentCreator(t)
	_ = dc.CreateDocument("one.md", "")
	_ = dc.CreateDocument("two.md", "")
	actual, err := tb.DocumentFileNames()
	if err != nil {
		t.Fatal(err)
	}
	expected := []domain.TikiDocumentFilename{
		domain.TikiDocumentFilename("one.md"),
		domain.TikiDocumentFilename("two.md"),
	}
	diff := cmp.Diff(expected, actual)
	if diff != "" {
		t.Fatal(diff)
	}
}

func TestDirectoryTikiBaseDocuments(t *testing.T) {
	tb1, dc := test.NewDocumentCreator(t)
	_ = dc.CreateDocument("one", "")
	_ = dc.CreateDocument("two", "")

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

func TestDirectoryTikiBaseLoad(t *testing.T) {
	tb := test.NewTempDirectoryTikiBase(t)
	expected, err := tb.CreateDocument("one.md", "The one")
	if err != nil {
		t.Fatal(err)
	}
	tb2, err := domain.NewDirectoryTikiBase(tb.StorageDir())
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
func TestDirectoryTikiBaseTikiLinks(t *testing.T) {
	tb, dc := test.NewDocumentCreator(t)
	doc1 := dc.CreateDocument("one.md", "# The one\n[The other](two.md)")
	doc2 := dc.CreateDocument("two.md", "# The other\n[The one](one.md)")
	actual, err := tb.TikiLinks()
	if err != nil {
		t.Fatal(err)
	}
	expected := []domain.TikiLink{
		domain.NewTikiLink("The other", doc1.TitleSection(), doc2),
		domain.NewTikiLink("The one", doc2.TitleSection(), doc1),
	}
	diff := cmp.Diff(expected, actual, cmp.AllowUnexported(expected[0], expected[0].SourceSection(), expected[0].TargetDocument()))
	if diff != "" {
		t.Fatal(diff)
	}
}

func TestNewDirectoryTikiBase(t *testing.T) {
	_, err := domain.NewDirectoryTikiBase(".")
	if err != nil {
		t.Fatal(err)
	}
}
