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

// newTikiDocument creates a new TikiDocument instance.
func newTikiDocument(filename TikiDocumentFilename, content string) TikiDocument {
	return TikiDocument{filename: filename, content: content}
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

func (td TikiDocument) TikiLinks(tb TikiBase) ([]TikiLink, error) {
	result := []TikiLink{}
	for _, section := range td.AllSections() {
		links, err := section.TikiLinks(tb)
		if err != nil {
			return result, err
		}
		result = append(result, links...)
	}
	return result, nil
}

// TikiSection provides the section before the content sections start.
func (td TikiDocument) TitleSection() TikiSection {
	return td.AllSections()[0]
}
