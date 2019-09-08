package domain

// TikiBase represents a collection of TikiDocuments.
type TikiBase interface {

	// CreateDocument creates a new TikiDocument with the given data in this DocumentCollection.
	CreateDocument(filename TikiDocumentFilename, content string) (TikiDocument, error)

	// DocumentHandles provides the handles for all documents in this collection,
	// sorted alphabetically.
	DocumentFileNames() ([]TikiDocumentFilename, error)

	// Documents provides all TikiDocuments in this collection.
	Documents() ([]TikiDocument, error)

	// Load provides the TikiDocument with the given filename, or an error if one doesn't exist.
	Load(filename TikiDocumentFilename) (TikiDocument, error)

	// TikiLinks provides all TikiLinks in all TikiDocuments within this TikiBase.
	TikiLinks() ([]TikiLink, error)
}
