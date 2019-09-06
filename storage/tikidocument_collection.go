package storage

import "fmt"

// TikiDocumentCollection represents an in-memory collection of TikiDocuments.
type TikiDocumentCollection struct {
	Documents []TikiDocument
}

// NewTikiDocumentCollection returns an empty TikiDocumentCollection.
func NewTikiDocumentCollection() *TikiDocumentCollection {
	return &TikiDocumentCollection{}
}

// Add adds the given TikiDocument to this document collection.
func (tdc *TikiDocumentCollection) Add(doc ...TikiDocument) {
	tdc.Documents = append(tdc.Documents, doc...)
}

// Find returns the TikiDocument from this collection with the given handle.
func (tdc *TikiDocumentCollection) Find(handle Handle) (TikiDocument, error) {
	for _, doc := range tdc.Documents {
		if doc.HasHandle(handle) {
			return doc, nil
		}
	}
	return TikiDocument{}, fmt.Errorf("document with handle '%s' not found", handle)
}
