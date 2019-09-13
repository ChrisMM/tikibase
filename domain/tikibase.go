package domain

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/pkg/errors"
)

// TikiBase represents a collection of TikiDocuments stored in a filesystem directory
// that form a knowledge base together.
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

// CreateDocument creates a new TikiDocument with the given content.
func (dtb TikiBase) CreateDocument(filename TikiDocumentFilename, content string) (result TikiDocument, err error) {
	doc := newTikiDocument(filename, content)
	filePath := path.Join(dtb.dir, string(doc.FileName()))
	err = ioutil.WriteFile(filePath, []byte(doc.content), 0644)
	return doc, err
}

// Documents returns all TikiDocuments in this TikiBase.
func (dtb TikiBase) Documents() (result TikiDocumentCollection, err error) {
	fileInfos, err := ioutil.ReadDir(dtb.dir)
	if err != nil {
		return result, errors.Wrap(err, "cannot read TikiBase directory")
	}
	for _, fileInfo := range fileInfos {
		doc, err := dtb.Load(TikiDocumentFilename(fileInfo.Name()))
		if err != nil {
			return result, errors.Wrapf(err, "cannot get all documents")
		}
		result = append(result, doc)
	}
	return result, nil
}

// Load provides the TikiDocument with the given filename, or an error if one doesn't exist.
func (dtb TikiBase) Load(filename TikiDocumentFilename) (result TikiDocument, err error) {
	path := path.Join(dtb.StorageDir(), string(filename))
	contentData, err := ioutil.ReadFile(path)
	if err != nil {
		return result, errors.Wrapf(err, "cannot read file '%s'", path)
	}
	return TikiDocument{filename, string(contentData)}, nil
}

// StorageDir provides the full directory path in which this TikiBase is stored.
func (dtb TikiBase) StorageDir() string {
	return dtb.dir
}
