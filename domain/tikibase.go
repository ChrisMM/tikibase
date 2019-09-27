package domain

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/pkg/errors"
)

// TikiBase represents a collection of Documents stored in a filesystem directory
// that form a knowledge base together.
type TikiBase struct {
	// the full path of the storage directory of this TikiBase
	dir string
}

// NewTikiBase creates a new TikiBase instance using the given directory path as its storage directory.
// The given file path must exist and be a directory.
func NewTikiBase(dir string) (result *TikiBase, err error) {
	info, err := os.Stat(dir)
	if err != nil {
		return result, err
	}
	if !info.IsDir() {
		return result, fmt.Errorf("%s is not a directory", dir)
	}
	return &TikiBase{dir}, nil
}

// CreateDocument creates a new Document with the given content.
func (tb *TikiBase) CreateDocument(filename DocumentFilename, content string) (result *Document, err error) {
	if !strings.HasSuffix(string(filename), ".md") {
		filename += ".md"
	}
	doc := newDocument(filename, content)
	filePath := path.Join(tb.dir, string(doc.FileName()))
	err = ioutil.WriteFile(filePath, []byte(doc.Content()), 0644)
	return doc, err
}

// Documents returns all Documents in this TikiBase.
func (tb *TikiBase) Documents() (result DocumentCollection, err error) {
	fileInfos, err := ioutil.ReadDir(tb.dir)
	if err != nil {
		return result, errors.Wrap(err, "cannot read TikiBase directory")
	}
	for i := range fileInfos {
		if !strings.HasSuffix(fileInfos[i].Name(), ".md") {
			continue
		}
		doc, err := tb.Load(DocumentFilename(fileInfos[i].Name()))
		if err != nil {
			return result, errors.Wrap(err, "cannot get all documents")
		}
		result = append(result, doc)
	}
	return result, nil
}

// Load provides the Document with the given filename, or an error if one doesn't exist.
func (tb *TikiBase) Load(filename DocumentFilename) (result *Document, err error) {
	if !strings.HasSuffix(string(filename), ".md") {
		filename += ".md"
	}
	path := path.Join(tb.StorageDir(), string(filename))
	contentData, err := ioutil.ReadFile(path)
	if err != nil {
		return result, errors.Wrapf(err, "cannot read file '%s'", path)
	}
	return newDocument(filename, string(contentData)), nil
}

// SaveDocument stores the given Document in this TikiBase.
func (tb *TikiBase) SaveDocument(doc *Document) error {
	filePath := path.Join(tb.dir, string(doc.FileName()))
	return ioutil.WriteFile(filePath, []byte(doc.Content()), 0644)
}

// StorageDir provides the full directory path in which this TikiBase is stored.
func (tb *TikiBase) StorageDir() string {
	return tb.dir
}
