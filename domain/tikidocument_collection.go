package domain

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

// DocumentHandles provides the handles of all documents stored in this TikiBase,
// sorted alphabetically.
func (tdc *TikiDocumentCollection) DocumentHandles() (result []Handle) {
	for _, doc := range tdc.Documents {
		result = append(result, doc.Handle())
	}
	return result
}

// Find returns the TikiDocument from this collection with the given handle.
func (tdc *TikiDocumentCollection) Find(handle Handle) (TikiDocument, error) {
	for _, doc := range tdc.Documents {
		if doc.Handle() == (handle) {
			return doc, nil
		}
	}
	return TikiDocument{}, fmt.Errorf("document with handle '%s' not found, existing documents are %v", handle, tdc.DocumentHandles())
}
