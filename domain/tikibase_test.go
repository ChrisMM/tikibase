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
	assert.Nil(t, err)
}

func TestTikiBase_CreateDocument(t *testing.T) {
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

func TestTikiBase_Documents(t *testing.T) {
	tikiBase := test.NewTempTikiBase(t)
	_, err := tikiBase.CreateDocument("one.md", "")
	assert.Nil(t, err)
	_, err = tikiBase.CreateDocument("two.md", "")
	assert.Nil(t, err)
	docs, err := tikiBase.Documents()
	assert.Nil(t, err)
	assert.Len(t, docs, 2)
	assert.Equal(t, domain.DocumentFilename("one.md"), docs[0].FileName())
	assert.Equal(t, domain.DocumentFilename("two.md"), docs[1].FileName())
}

func TestTikiBase_Documents_IgnoresNonMarkdown(t *testing.T) {
	tikiBase := test.NewTempTikiBase(t)
	_, err := tikiBase.CreateDocument("one.md", "")
	assert.Nil(t, err)
	err = test.CreateBinaryFile(path.Join(tikiBase.StorageDir(), "foo.png"))
	assert.Nil(t, err)
	docs, err := tikiBase.Documents()
	assert.Nil(t, err, "cannot call tb.Documents()")
	assert.Len(t, docs, 1)
	assert.Equal(t, domain.DocumentFilename("one.md"), docs[0].FileName())
}

func TestTikiBase_Load(t *testing.T) {
	tikiBase := test.NewTempTikiBase(t)
	_, err := tikiBase.CreateDocument("one.md", "The one")
	assert.Nil(t, err)
	doc, err := tikiBase.LoadDocument("one.md")
	assert.Nil(t, err)
	assert.Equal(t, domain.DocumentFilename("one.md"), doc.FileName())
}

func TestTikiBase_SaveDocument(t *testing.T) {
	tikiBase := test.NewTempTikiBase(t)
	doc := domain.ScaffoldDocument(domain.DocumentScaffold{FileName: "one.md", Content: "document content"})
	err := tikiBase.SaveDocument(doc)
	assert.Nil(t, err, "cannot save document")
	filePath := path.Join(tikiBase.StorageDir(), "one.md")
	fileInfo, err := os.Stat(filePath)
	assert.Nil(t, err, "file not found")
	assert.False(t, fileInfo.IsDir(), "file should not be a directory")
	fileMode := fileInfo.Mode()
	assert.Equalf(t, os.FileMode(0644), fileMode, "file should have access 0644 but has %#o", fileMode)
}

func TestTikiBase_StorageDir(t *testing.T) {
	currentDir, err := os.Getwd()
	assert.Nil(t, err, "cannot determine current working directory")
	tb, err := domain.NewTikiBase(currentDir)
	assert.Nil(t, err)
	actual := tb.StorageDir()
	assert.Equal(t, currentDir, actual)
}
