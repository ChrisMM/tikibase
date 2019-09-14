package domain

import (
	"strings"

	"github.com/kevgo/tikibase/helpers"
)

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
func (td Document) AllSections() (result SectionCollection) {
	tsb := NewSectionBuilder("", &td)
	lines := helpers.CutStringIntoLines(td.content)
	for _, line := range lines {
		if strings.HasPrefix(line, "#") {
			if tsb.Len() > 0 {
				result = append(result, tsb.Section())
			}
			tsb = NewSectionBuilder(line, &td)
		} else {
			tsb.AddLine(line)
		}
	}
	result = append(result, tsb.Section())
	return result
}

// AppendSection provides a new Document with the given Section appended.
func (td Document) AppendSection(section Section) Document {
	// add an empty line to the last section
	sections := td.AllSections()
	lastSection := sections[len(sections)-1]
	newLastSection := lastSection.AppendLine("\n")
	sections = sections.Replace(lastSection, newLastSection)

	// add the new section
	newSections := append(sections, section)
	return newDocument(td.filename, newSections.Text())
}

// Content returns the content of this document.
func (td Document) Content() string {
	return td.content
}

// ContentSections provides the content sections of this document.
func (td Document) ContentSections() SectionCollection {
	return td.AllSections()[1:]
}

// FileName returns the file path (handle + extension) of this Document.
func (td Document) FileName() DocumentFilename {
	return td.filename
}

// ReplaceSection provides a new Document that is like this one
// and has the given old section replaced with the given new section.
func (td Document) ReplaceSection(oldSection, newSection Section) Document {
	newSections := td.AllSections().Replace(oldSection, newSection)
	return newDocument(td.filename, newSections.Text())
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

// URL provides the URL of this Document within its TikiBase.
func (td Document) URL() string {
	return string(td.filename)
}
