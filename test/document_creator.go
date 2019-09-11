package test

import (
	"testing"

	"github.com/kevgo/tikibase/domain"
)

// DocumentCreator makes creating lots of TikiDocuments for testing easy.
// Errors during dc.CreateDocument() cause the test to fail automatically.
type DocumentCreator struct {
	t   *testing.T
	tb  domain.TikiBase
	err error
}

// NewDocumentCreator provides a DocumentCreator instance operating in the system's temp directory.
func NewDocumentCreator(t *testing.T) (domain.DirectoryTikiBase, DocumentCreator) {
	tb := NewTempDirectoryTikiBase(t)
	return tb, DocumentCreator{t, tb, nil}
}

// CreateDocument creates and provides a new TikiDocument.
// It has the same API as TikiBase.CreateDocument except it doesn't return errors.
// You have to check for errors after you are done by calling Err().
func (dc *DocumentCreator) CreateDocument(filename domain.TikiDocumentFilename, content string) domain.TikiDocument {
	result, err := dc.tb.CreateDocument(filename, content)
	if err != nil {
		dc.t.Fatalf("error creating document '%s': %v", filename, err)
	}
	return result
}
