package domain_test

import (
	"os"
	"path"
	"testing"

	"github.com/kevgo/tikibase/domain"
	"github.com/kevgo/tikibase/test"
	"github.com/stretchr/testify/assert"
)

func TestNewTikiBase(t *testing.T) {
	_, err := domain.NewTikiBase(".")
	if err != nil {
		t.Fatal(err)
	}
}

func TestTikiBaseCreateDocument(t *testing.T) {
	tb := test.NewTempTikiBase(t)
	_, err := tb.CreateDocument("one", "The one.")
	assert.Nil(t, err, "cannot create document")
	filePath := path.Join(tb.StorageDir(), "one.md")
	fileInfo, err := os.Stat(filePath)
	assert.Nil(t, err, "file not found:", filePath)
	assert.False(t, fileInfo.IsDir(), "file should not be a directory")
	fileMode := fileInfo.Mode()
	assert.Equalf(t, os.FileMode(0644), fileMode, "file should have access 0644 but has %#o", fileMode)
}

func TestTikiBaseDocumentsIgnoresNonMarkdown(t *testing.T) {
	tb1 := test.NewTempTikiBase(t)
	_, err := tb1.CreateDocument("one", "")
	assert.Nil(t, err)
	err = test.CreateBinaryFile(path.Join(tb1.StorageDir(), "foo.png"))
	assert.Nil(t, err)

	// get the documents
	tb2, err := domain.NewTikiBase(tb1.StorageDir())
	assert.Nil(t, err, "cannot instantiate tb2")
	actual, err := tb2.Documents()
	assert.Nil(t, err, "cannot call tb.Documents()")

	// verify results
	assert.Len(t, actual, 1)
}

func TestTikiBaseDocuments(t *testing.T) {
	tb1 := test.NewTempTikiBase(t)
	_, err := tb1.CreateDocument("one", "")
	assert.Nil(t, err)
	_, err = tb1.CreateDocument("two", "")
	assert.Nil(t, err)

	// get the documents
	tb2, err := domain.NewTikiBase(tb1.StorageDir())
	assert.Nil(t, err, "cannot instantiate tb2")
	actual, err := tb2.Documents()
	assert.Nil(t, err, "cannot call tb.Documents()")

	// verify results
	assert.Len(t, actual, 2)
}

func TestTikiBaseLoad(t *testing.T) {
	tb := test.NewTempTikiBase(t)
	expected, err := tb.CreateDocument("one.md", "The one")
	assert.Nil(t, err)
	tb2, err := domain.NewTikiBase(tb.StorageDir())
	assert.Nil(t, err)
	actual, err := tb2.Load("one.md")
	assert.Nil(t, err)
	assert.Equal(t, expected.Content(), actual.Content(), "mismatching content")
}

func TestTikiBaseSaveDocument(t *testing.T) {
	tb := test.NewTempTikiBase(t)
	doc := domain.ScaffoldDocument(domain.DocumentScaffold{FileName: "one.md", Content: "document content"})
	err := tb.SaveDocument(doc)
	assert.Nil(t, err, "cannot save document")
	filePath := path.Join(tb.StorageDir(), "one.md")
	fileInfo, err := os.Stat(filePath)
	assert.Nil(t, err, "file not found")
	assert.False(t, fileInfo.IsDir(), "file should not be a directory")
	fileMode := fileInfo.Mode()
	assert.Equalf(t, os.FileMode(0644), fileMode, "file should have access 0644 but has %#o", fileMode)
}

func TestTikiBaseStorageDir(t *testing.T) {
	currentDir, err := os.Getwd()
	assert.Nil(t, err, "cannot determine current working directory")
	tb, err := domain.NewTikiBase(currentDir)
	assert.Nil(t, err)
	actual := tb.StorageDir()
	assert.Equal(t, currentDir, actual, "wrong StorageDir provided")
}
