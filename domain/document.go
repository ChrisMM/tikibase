package domain

import "strings"

// Document represents a MarkDown file in the document base.
// Create new instances via DocumentCollection.CreateDocument
type Document struct {

	// fileame is the unique identifier for this document.
	filename DocumentFilename

	// the textual content of the document
	content string
}

// DocumentFilename is the filename of a Document.
type DocumentFilename string

// DocumentScaffold is for easy scaffolding of Documents in tests.
// Don't use this in production code.
type DocumentScaffold struct {
	FileName string
	Content  string
}

// newDocument creates a new Document instance.
// This constructor is internal to this module,
// call (TikiBase).CreateDocument() to create new documents in production.
func newDocument(filename DocumentFilename, content string) Document {
	return Document{filename: filename, content: content}
}

// ScaffoldDocument provides new Documents for testing.
func ScaffoldDocument(data DocumentScaffold) Document {
	if data.FileName == "" {
		data.FileName = "default.md"
	}
	if data.Content == "" {
		data.Content = "# Title\ndefault content"
	}
	return newDocument(DocumentFilename(data.FileName), data.Content)
}

// AllSections returns all the TikiSections that make up this document,
// including the title section.
func (td Document) AllSections() []Section {
	result := []Section{}
	var tsb SectionBuilder
	for i, line := range strings.Split(td.content, "\n") {
		if strings.HasPrefix(line, "#") {
			if i > 0 {
				result = append(result, tsb.Section())
			}
			tsb = NewSectionBuilder(line)
		} else {
			tsb.AddLine(line)
		}
	}
	result = append(result, tsb.Section())
	return result
}

// FileName returns the file path (handle + extension) of this Document.
func (td Document) FileName() DocumentFilename {
	return td.filename
}

// TikiLinks returns the TikiLinks in this Document.
func (td Document) TikiLinks(tdc DocumentCollection) (result TikiLinkCollection, err error) {
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
func (td Document) TitleSection() Section {
	return td.AllSections()[0]
}
