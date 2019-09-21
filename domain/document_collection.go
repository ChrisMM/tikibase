package domain

import (
	"fmt"
)

// DocumentCollection is a collection of Documents.
type DocumentCollection []Document

// ScaffoldDocumentCollection provides new DocumentCollections for testing.
func ScaffoldDocumentCollection(data []DocumentScaffold) (result DocumentCollection) {
	for i := range data {
		result = append(result, ScaffoldDocument(data[i]))
	}
	return result
}

// FileNames returns the filenames of all documents in this DocumentCollection.
func (dc DocumentCollection) FileNames() (result []DocumentFilename, err error) {
	for i := range dc {
		result = append(result, dc[i].FileName())
	}
	return result, nil
}

// Find returns the Document with the given filename.
func (dc DocumentCollection) Find(filename DocumentFilename) (result *Document, err error) {
	for i := range dc {
		if dc[i].FileName() == filename {
			return &dc[i], nil
		}
	}
	return result, fmt.Errorf("cannot find document '%s'", filename)
}

// TikiLinks provides all TikiLinks in all Documents within this TikiBase.
func (dc DocumentCollection) TikiLinks() (result TikiLinkCollection, err error) {
	for i := range dc {
		links, err := dc[i].TikiLinks(dc)
		if err != nil {
			return result, err
		}
		result = append(result, links...)
	}
	return result, nil
}
