package domain

import (
	"fmt"

	"github.com/pkg/errors"
)

// DocumentCollection is a collection of Documents.
type DocumentCollection []Document

// ScaffoldDocumentCollection provides new DocumentCollections for testing.
func ScaffoldDocumentCollection(data []DocumentScaffold) (result DocumentCollection) {
	for _, docData := range data {
		result = append(result, ScaffoldDocument(docData))
	}
	return result
}

// FileNames returns the filenames of all documents in this DocumentCollection.
func (tdc DocumentCollection) FileNames() (result []DocumentFilename, err error) {
	for _, doc := range tdc {
		result = append(result, doc.FileName())
	}
	return result, nil
}

// Find returns the Document with the given filename.
func (tdc DocumentCollection) Find(filename DocumentFilename) (result Document, err error) {
	for _, doc := range tdc {
		if doc.FileName() == filename {
			return doc, nil
		}
	}
	return result, fmt.Errorf("cannot find document '%s'", filename)
}

// TikiLinks provides all TikiLinks in all Documents within this TikiBase.
func (tdc DocumentCollection) TikiLinks() (result TikiLinkCollection, err error) {
	for _, doc := range tdc {
		links, err := doc.TikiLinks(tdc)
		if err != nil {
			return result, errors.Wrapf(err, "cannot determine the TikiLinks of '%+v'", doc)
		}
		result = append(result, links...)
	}
	return result, nil
}
