package domain

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/pkg/errors"
)

// DirectoryTikiBase represents a collection of TikiDocuments stored in a filesystem directory
// that form a knowledge base together.
type DirectoryTikiBase struct {
	// the full path of the storage directory of this TikiBase
	dir string
}

// NewDirectoryTikiBase creates a new TikiBase instance using the given directory path as its storage directory.
// The given file path must exist and be a directory.
func NewDirectoryTikiBase(dir string) (result DirectoryTikiBase, err error) {
	info, err := os.Stat(dir)
	if err != nil {
		return result, err
	}
	if !info.IsDir() {
		return result, fmt.Errorf("%s is not a directory", dir)
	}
	return DirectoryTikiBase{dir}, nil
}

// CreateDocument creates a new TikiDocument with the given content.
func (dtb DirectoryTikiBase) CreateDocument(filename TikiDocumentFilename, content string) (result TikiDocument, err error) {
	doc := newTikiDocument(filename, content)
	filePath := path.Join(dtb.dir, string(doc.FileName()))
	err = ioutil.WriteFile(filePath, []byte(doc.content), 0644)
	return doc, err
}

// DocumentFileNames returns the filenames of all documents in this DirectoryTikiBase.
func (dtb DirectoryTikiBase) DocumentFileNames() (result []TikiDocumentFilename, err error) {
	docs, err := dtb.Documents()
	if err != nil {
		return result, err
	}
	for _, doc := range docs {
		result = append(result, doc.FileName())
	}
	return result, nil
}

// Documents returns all TikiDocuments in this TikiBase.
func (dtb DirectoryTikiBase) Documents() (result []TikiDocument, err error) {
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
func (dtb DirectoryTikiBase) Load(filename TikiDocumentFilename) (result TikiDocument, err error) {
	path := path.Join(dtb.StorageDir(), string(filename))
	contentData, err := ioutil.ReadFile(path)
	if err != nil {
		return result, errors.Wrapf(err, "cannot read file '%s'", path)
	}
	return TikiDocument{filename, string(contentData)}, nil
}

// StorageDir provides the full directory path in which this TikiBase is stored.
func (dtb DirectoryTikiBase) StorageDir() string {
	return dtb.dir
}

// TikiLinks provides all TikiLinks in all TikiDocuments within this TikiBase.
func (dtb DirectoryTikiBase) TikiLinks() (result []TikiLink, err error) {
	docs, err := dtb.Documents()
	if err != nil {
		return result, errors.Wrapf(err, "cannot determine the TikiLinks of '%+v'", dtb)
	}
	for _, doc := range docs {
		links, err := doc.TikiLinks(dtb)
		if err != nil {
			return result, errors.Wrapf(err, "cannot determine the TikiLinks of '%+v'", doc)
		}
		result = append(result, links...)
	}
	return result, nil
}
