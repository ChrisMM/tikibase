package domain

import (
	"fmt"

	"github.com/kevgo/tikibase/helpers"
)

// Documents is a collection of Documents.
type Documents []*Document

// ScaffoldDocuments provides new DocumentCollections for testing.
func ScaffoldDocuments(data []DocumentScaffold) (result Documents) {
	for i := range data {
		result = append(result, ScaffoldDocument(data[i]))
	}
	return result
}

// Contains indicates whether this DocumentCollection contains the given document.
func (docs Documents) Contains(doc *Document) bool {
	for i := range docs {
		if docs[i] == doc {
			return true
		}
	}
	return false
}

// ContentSections provides all content sections in this document collection.
func (docs Documents) ContentSections() (result Sections) {
	for d := range docs {
		result = append(result, docs[d].ContentSections()...)
	}
	return result
}

// Links provides the links in this document collection.
func (docs Documents) Links() (internal, external Links) {
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
func (docs Documents) FindByFilename(filename string) (result *Document, err error) {
	for i := range docs {
		if docs[i].FileName() == filename {
			return docs[i], nil
		}
	}
	return result, fmt.Errorf("cannot find document %q", filename)
}

// TikiLinks provides all TikiLinks in all Documents within this TikiBase.
func (docs Documents) TikiLinks() (result TikiLinks, err error) {
	for i := range docs {
		links, err := docs[i].TikiLinks(docs)
		if err != nil {
			return result, err
		}
		result = append(result, links...)
	}
	return result, nil
}
