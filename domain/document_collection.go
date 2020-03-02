package domain

import (
	"fmt"
)

// DocumentCollection is a collection of Documents.
type DocumentCollection []*Document

// ScaffoldDocumentCollection provides new DocumentCollections for testing.
func ScaffoldDocumentCollection(data []DocumentScaffold) (result DocumentCollection) {
	for i := range data {
		result = append(result, ScaffoldDocument(data[i]))
	}
	return result
}

// Contains indicates whether this DocumentCollection contains the given document.
func (docs DocumentCollection) Contains(doc *Document) bool {
	for i := range docs {
		if docs[i] == doc {
			return true
		}
	}
	return false
}

// FindByFilename returns the Document with the given filename.
func (docs DocumentCollection) FindByFilename(filename DocumentFilename) (result *Document, err error) {
	for i := range docs {
		if docs[i].FileName() == filename {
			return docs[i], nil
		}
	}
	return result, fmt.Errorf("cannot find document %q", filename)
}

// TikiLinks provides all TikiLinks in all Documents within this TikiBase.
func (docs DocumentCollection) TikiLinks() (result TikiLinkCollection, err error) {
	for i := range docs {
		links, err := docs[i].TikiLinks(docs)
		if err != nil {
			return result, err
		}
		result = append(result, links...)
	}
	return result, nil
}
