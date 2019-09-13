package domain

import "strings"

// TikiDocument represents a MarkDown file in the document base.
// Create new instances via TikiDocumentCollection.CreateDocument
type TikiDocument struct {

	// fileame is the unique identifier for this document.
	filename TikiDocumentFilename

	// the textual content of the document
	content string
}

// TikiDocumentScaffold is for easy scaffolding of TikiDocuments in tests.
// Don't use this in production code.
type TikiDocumentScaffold struct {
	FileName string
	Content  string
}

// newTikiDocument creates a new TikiDocument instance.
// This constructor is internal to this module,
// call (TikiBase).CreateDocument() to create new documents in production.
func newTikiDocument(filename TikiDocumentFilename, content string) TikiDocument {
	return TikiDocument{filename: filename, content: content}
}

// ScaffoldTikiDocument scaffolds a new TikiDocument.
// This method is just for testing, don't use it in production code.
func ScaffoldTikiDocument(data TikiDocumentScaffold) TikiDocument {
	if data.FileName == "" {
		data.FileName = "default.md"
	}
	if data.Content == "" {
		data.Content = "default content"
	}
	return newTikiDocument(TikiDocumentFilename(data.FileName), data.Content)
}

// AllSections returns all the TikiSections that make up this document,
// including the title section.
func (td TikiDocument) AllSections() []TikiSection {
	result := []TikiSection{}
	var tsb TikiSectionBuilder
	for i, line := range strings.Split(td.content, "\n") {
		if strings.HasPrefix(line, "#") {
			if i > 0 {
				result = append(result, tsb.Section())
			}
			tsb = NewTikiSectionBuilder(line)
		} else {
			tsb.AddLine(line)
		}
	}
	result = append(result, tsb.Section())
	return result
}

// FileName returns the file path (handle + extension) of this TikiDocument.
func (td TikiDocument) FileName() TikiDocumentFilename {
	return td.filename
}

// TikiLinks returns the TikiLinks in this TikiDocument.
func (td TikiDocument) TikiLinks(tdc TikiDocumentCollection) ([]TikiLink, error) {
	result := []TikiLink{}
	for _, section := range td.AllSections() {
		links, err := section.TikiLinks(tdc)
		if err != nil {
			return result, err
		}
		result = append(result, links...)
	}
	return result, nil
}

// TitleSection provides the section before the content sections start.
func (td TikiDocument) TitleSection() TikiSection {
	return td.AllSections()[0]
}
