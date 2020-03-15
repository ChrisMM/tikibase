package domain

import (
	"fmt"

	"github.com/kevgo/tikibase/helpers"
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

// Links provides the links in this document collection.
func (docs DocumentCollection) Links() (internal, external LinkCollection) {
	for d := range docs {
		links := docs[d].Links()
		for l := range links {
			if helpers.IsURL(links[l].Target()) {
				external = append(external, links[l])
			} else {
				internal = append(internal, links[l])
			}
		}
	}
	return internal, external
}

// FindByFilename returns the Document with the given filename.
func (docs DocumentCollection) FindByFilename(filename string) (result *Document, err error) {
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
