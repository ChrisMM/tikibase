package domain_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/kevgo/tikibase/config"
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
	filePath := filepath.Join(tb.StorageDir(), "one.md")
	fileInfo, err := os.Stat(filePath)
	assert.Nil(t, err, "file not found:", filePath)
	assert.False(t, fileInfo.IsDir(), "file should not be a directory")
	fileMode := fileInfo.Mode()
	assert.Equalf(t, os.FileMode(domain.NewFilePerms), fileMode, "file should have access %o but has %#o", domain.NewFilePerms, fileMode)
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
	assert.Equal(t, "one.md", docs[0].FileName())
	assert.Equal(t, "two.md", docs[1].FileName())
}

func TestTikiBase_Documents_IgnoresNonMarkdown(t *testing.T) {
	tikiBase := test.NewTempTikiBase(t)
	_, err := tikiBase.CreateDocument("one.md", "")
	assert.Nil(t, err)
	err = test.CreateBinaryFile(filepath.Join(tikiBase.StorageDir(), "foo.png"))
	assert.Nil(t, err)
	docs, err := tikiBase.Documents()
	assert.Nil(t, err, "cannot call tb.Documents()")
	assert.Len(t, docs, 1)
	assert.Equal(t, "one.md", docs[0].FileName())
}

func TestTikiBase_Files_noConfig(t *testing.T) {
	tmp, err := ioutil.TempDir("", "")
	assert.Nil(t, err)
	_, err = os.Create(filepath.Join(tmp, "one.md"))
	assert.Nil(t, err)
	_, err = os.Create(filepath.Join(tmp, "two.md"))
	assert.Nil(t, err)
	_, err = os.Create(filepath.Join(tmp, "img.png"))
	assert.Nil(t, err)
	tb, err := domain.NewTikiBase(tmp)
	assert.Nil(t, err)
	docs, resources, err := tb.Files()
	assert.Nil(t, err)
	assert.Equal(t, []string{"one.md", "two.md"}, docs.FileNames())
	assert.Equal(t, []string{"img.png"}, []string(resources))
}

func TestTikiBase_Files_withConfig(t *testing.T) {
	dir, err := ioutil.TempDir("", "")
	assert.Nil(t, err)
	_, err = os.Create(filepath.Join(dir, "one.md"))
	assert.Nil(t, err)
	_, err = os.Create(filepath.Join(dir, "Makefile"))
	assert.Nil(t, err)
	_, err = os.Create(filepath.Join(dir, "img.png"))
	assert.Nil(t, err)
	err = ioutil.WriteFile(filepath.Join(dir, config.FileName()), []byte("ignore:\n  - Makefile\n"), 0644)
	assert.Nil(t, err)
	tb, err := domain.NewTikiBase(dir)
	assert.Nil(t, err)
	fmt.Println(tb)
	docs, resources, err := tb.Files()
	assert.Nil(t, err)
	assert.Equal(t, []string{"one.md"}, docs.FileNames())
	assert.Equal(t, []string{"img.png"}, []string(resources))
}

func TestTikiBase_Load(t *testing.T) {
	tikiBase := test.NewTempTikiBase(t)
	_, err := tikiBase.CreateDocument("one.md", "The one")
	assert.Nil(t, err)
	doc, err := tikiBase.LoadDocument("one.md")
	assert.Nil(t, err)
	assert.Equal(t, "one.md", doc.FileName())
}

func TestTikiBase_SaveDocument(t *testing.T) {
	tikiBase := test.NewTempTikiBase(t)
	doc := domain.ScaffoldDocument(domain.DocumentScaffold{FileName: "one.md", Content: "document content"})
	err := tikiBase.SaveDocument(doc)
	assert.Nil(t, err, "cannot save document")
	filePath := filepath.Join(tikiBase.StorageDir(), "one.md")
	fileInfo, err := os.Stat(filePath)
	assert.Nil(t, err, "file not found")
	assert.False(t, fileInfo.IsDir(), "file should not be a directory")
	fileMode := fileInfo.Mode()
	assert.Equalf(t, os.FileMode(domain.NewFilePerms), fileMode, "file should have access 0644 but has %#o", fileMode)
}

func TestTikiBase_StorageDir(t *testing.T) {
	currentDir, err := os.Getwd()
	assert.Nil(t, err, "cannot determine current working directory")
	tb, err := domain.NewTikiBase(currentDir)
	assert.Nil(t, err)
	actual := tb.StorageDir()
	assert.Equal(t, currentDir, actual)
}
