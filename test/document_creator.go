package test

import (
	"testing"

	"github.com/kevgo/tikibase/domain"
)

// DocumentCreator makes creating lots of TikiDocuments easy.
// It emulates the API of [bufio.Scanner](https://golang.org/pkg/bufio/#Scanner):
// You can call dc.CreateDocument() multiple times without checking for errors.
// Errors are accumulated internally and checked at the end by calling dc.Err().
type DocumentCreator struct {
	tb  domain.TikiBase
	err error
}

// NewDocumentCreator provides a DocumentCreator instance for the given TikiBase.
func NewDocumentCreator(tb domain.TikiBase) *DocumentCreator {
	return &DocumentCreator{tb, nil}
}

// NewTestDocumentCreator provides a DocumentCreator instance operating in the system's temp directory.
func NewTestDocumentCreator(t *testing.T) (domain.DirectoryTikiBase, *DocumentCreator) {
	tb := NewTempDirectoryTikiBase(t)
	return tb, NewDocumentCreator(tb)
}

// CreateDocument creates and provides a new TikiDocument.
// It has the same API as TikiBase.CreateDocument except it doesn't return errors.
// You have to check for errors after you are done by calling Err().
func (dc *DocumentCreator) CreateDocument(filename domain.TikiDocumentFilename, content string) domain.TikiDocument {
	result, err := dc.tb.CreateDocument(filename, content)
	if err != nil {
		dc.err = err
	}
	return result
}

// Err provides any error encountered while creating TikiDocuments before.
func (dc *DocumentCreator) Err() error {
	return dc.err
}
