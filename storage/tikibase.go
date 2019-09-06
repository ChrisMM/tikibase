package storage

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/pkg/errors"
)

// TikiBase represents a collection of Markdown files that form a knowledge base together.
type TikiBase struct {
	// the full path of the storage directory of this TikiBase
	dir string
}

// NewTikiBase creates a new TikiBase instance using the given directory path as its storage directory.
// The given file path must exist and be a directory.
func NewTikiBase(dir string) (result TikiBase, err error) {
	info, err := os.Stat(dir)
	if err != nil {
		return result, err
	}
	if !info.IsDir() {
		return result, fmt.Errorf("%s is not a directory", dir)
	}
	return TikiBase{dir}, nil
}

// CreateDocument creates a new TikiDocument in this TikiBase and returns it.
func (tb TikiBase) CreateDocument(handle Handle, content string) (result TikiDocument, err error) {
	doc := NewTikiDocument(handle, content)
	return doc, tb.SaveDocument(doc)
}

// Documents returns all TikiDocuments in this TikiBase.
func (tb TikiBase) Documents() (result []TikiDocument, err error) {
	fileInfos, err := ioutil.ReadDir(tb.dir)
	if err != nil {
		return result, errors.Wrap(err, "cannot read TikiBase directory")
	}
	for _, fileInfo := range fileInfos {
		path := path.Join(tb.StorageDir(), fileInfo.Name())
		contentData, err := ioutil.ReadFile(path)
		if err != nil {
			return result, errors.Wrapf(err, "cannot read file '%s'", path)
		}
		result = append(result, NewTikiDocument(NewHandleFromFileName(path), string(contentData)))
	}
	return result, nil
}

// SaveDocument persists the given TikiDocument into this TikiBase
func (tb TikiBase) SaveDocument(doc TikiDocument) error {
	return ioutil.WriteFile(path.Join(tb.dir, doc.FilePath()), []byte(doc.content), 0644)
}

// StorageDir provides the full directory path in which this TikiBase is stored.
func (tb TikiBase) StorageDir() string {
	return tb.dir
}
