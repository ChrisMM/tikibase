package domain

import (
	"fmt"

	"github.com/pkg/errors"
)

// TikiDocumentCollection is a collection of TikiDocuments.
type TikiDocumentCollection []TikiDocument

// ScaffoldTikiDocumentCollection creates a new TikiDocumentCollection with the given data.
// This is only for tests, don't use this in production code.
func ScaffoldTikiDocumentCollection(data []TikiDocumentScaffold) (result TikiDocumentCollection) {
	for _, docData := range data {
		result = append(result, ScaffoldTikiDocument(docData))
	}
	return result
}

// FileNames returns the filenames of all documents in this TikiDocumentCollection.
func (tdc TikiDocumentCollection) FileNames() (result []TikiDocumentFilename, err error) {
	for _, doc := range tdc {
		result = append(result, doc.FileName())
	}
	return result, nil
}

// Find returns the TikiDocument with the given filename.
func (tdc TikiDocumentCollection) Find(filename TikiDocumentFilename) (result TikiDocument, err error) {
	for _, doc := range tdc {
		if doc.FileName() == filename {
			return doc, nil
		}
	}
	return result, fmt.Errorf("Cannot find document '%s'", filename)
}

// TikiLinks provides all TikiLinks in all TikiDocuments within this TikiBase.
func (tdc TikiDocumentCollection) TikiLinks() (result TikiLinkCollection, err error) {
	for _, doc := range tdc {
		links, err := doc.TikiLinks(tdc)
		if err != nil {
			return result, errors.Wrapf(err, "cannot determine the TikiLinks of '%+v'", doc)
		}
		result = append(result, links...)
	}
	return result, nil
}
