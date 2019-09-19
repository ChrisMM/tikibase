package domain

import (
	"fmt"

	"github.com/pkg/errors"
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
		fmt.Printf("tdc.TikiLinks(): doc = %p, %s\n", &dc[i], dc[i].TitleSection().Anchor())
		links, err := dc[i].TikiLinks(dc)
		if err != nil {
			return result, errors.Wrapf(err, "cannot determine the TikiLinks of '%+v'", dc[i])
		}
		result = append(result, links...)
	}
	return result, nil
}
