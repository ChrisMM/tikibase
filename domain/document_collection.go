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
func (tdc DocumentCollection) FileNames() (result []DocumentFilename, err error) {
	for i := range tdc {
		result = append(result, tdc[i].FileName())
	}
	return result, nil
}

// Find returns the Document with the given filename.
func (tdc DocumentCollection) Find(filename DocumentFilename) (result *Document, err error) {
	for i := range tdc {
		if tdc[i].FileName() == filename {
			return &tdc[i], nil
		}
	}
	return result, fmt.Errorf("cannot find document '%s'", filename)
}

// TikiLinks provides all TikiLinks in all Documents within this TikiBase.
func (tdc DocumentCollection) TikiLinks() (result TikiLinkCollection, err error) {
	for i := range tdc {
		fmt.Printf("tdc.TikiLinks(): doc = %p, %s\n", &tdc[i], tdc[i].TitleSection().Anchor())
		links, err := tdc[i].TikiLinks(tdc)
		if err != nil {
			return result, errors.Wrapf(err, "cannot determine the TikiLinks of '%+v'", tdc[i])
		}
		result = append(result, links...)
	}
	return result, nil
}
