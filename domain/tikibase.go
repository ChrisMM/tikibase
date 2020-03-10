package domain

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
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
		return result, fmt.Errorf("cannot create TikiBase in directory %q: %w", dir, err)
	}
	if !info.IsDir() {
		return result, fmt.Errorf("%s must be a directory to contain a TikiBase", dir)
	}
	return &TikiBase{dir}, nil
}

// CreateDocument creates a new Document with the given content.
func (tikiBase *TikiBase) CreateDocument(filename DocumentFilename, content string) (result *Document, err error) {
	if !strings.HasSuffix(string(filename), ".md") {
		filename += ".md"
	}
	doc := newDocumentWithText(filename, content)
	filePath := path.Join(tikiBase.dir, string(doc.FileName()))
	err = ioutil.WriteFile(filePath, []byte(doc.Content()), 0644)
	if err != nil {
		return result, fmt.Errorf("cannot create new document %q: %w", filename, err)
	}
	return doc, nil
}

// Documents returns all Documents in this TikiBase.
func (tikiBase *TikiBase) Documents() (result DocumentCollection, err error) {
	fileInfos, err := ioutil.ReadDir(tikiBase.dir)
	if err != nil {
		return result, fmt.Errorf("cannot read TikiBase directory %q: %w", tikiBase.dir, err)
	}
	for i := range fileInfos {
		if !strings.HasSuffix(fileInfos[i].Name(), ".md") {
			continue
		}
		doc, err := tikiBase.LoadDocument(DocumentFilename(fileInfos[i].Name()))
		if err != nil {
			return result, fmt.Errorf("cannot get all documents in TikiBase %q: %w", tikiBase.dir, err)
		}
		result = append(result, doc)
	}
	return result, nil
}

// LoadDocument provides the Document with the given filename, or an error if one doesn't exist.
func (tikiBase *TikiBase) LoadDocument(filename DocumentFilename) (result *Document, err error) {
	if !strings.HasSuffix(string(filename), ".md") {
		filename += ".md"
	}
	path := path.Join(tikiBase.StorageDir(), string(filename))
	contentData, err := ioutil.ReadFile(path)
	if err != nil {
		return result, fmt.Errorf("cannot load TikiBase document %q: %w", path, err)
	}
	return newDocumentWithText(filename, string(contentData)), nil
}

// SaveDocument stores the given Document in this TikiBase.
func (tikiBase *TikiBase) SaveDocument(doc *Document) error {
	filePath := path.Join(tikiBase.dir, string(doc.FileName()))
	err := ioutil.WriteFile(filePath, []byte(doc.Content()), 0644)
	if err != nil {
		return fmt.Errorf("cannot save document %q: %w", doc.FileName(), err)
	}
	return nil
}

// StorageDir provides the full directory path in which this TikiBase is stored.
func (tikiBase *TikiBase) StorageDir() string {
	return tikiBase.dir
}
