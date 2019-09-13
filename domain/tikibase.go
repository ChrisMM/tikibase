package domain

// TikiBase represents a collection of TikiDocuments.
type TikiBase interface {

	// CreateDocument creates a new TikiDocument with the given data in this DocumentCollection.
	CreateDocument(filename TikiDocumentFilename, content string) (TikiDocument, error)

	// Documents provides all TikiDocuments in this collection.
	Documents() (TikiDocumentCollection, error)

	// Load provides the TikiDocument with the given filename, or an error if one doesn't exist.
	Load(filename TikiDocumentFilename) (TikiDocument, error)
}
