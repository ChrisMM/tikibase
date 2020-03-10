package domain

import (
	"fmt"
	"path"
	"strings"
)

// Document represents a MarkDown file in the document base.
// Create new instances via DocumentCollection.CreateDocument
type Document struct {

	// fileame is the unique identifier for this document.
	filename DocumentFilename

	// the textual content of the document
	sections SectionCollection
}

// DocumentFilename is the filename of a Document.
type DocumentFilename string

// DocumentScaffold is for easy scaffolding of Documents in tests.
// Don't use this in production code.
type DocumentScaffold struct {
	FileName string
	Content  string
}

// newDocumentWithText creates a new Document instance with the given textual content.
// This constructor is internal to this module,
// call (TikiBase).CreateDocument() to create new documents in production.
func newDocumentWithText(filename DocumentFilename, content string) *Document {
	doc := Document{filename: filename}
	doc.sections = newSectionCollection(content, &doc)
	return &doc
}

// newDocumentWithSections creates a new Document instance with the given pre-parsed sections.
// This constructor is internal to this module,
// call (TikiBase).CreateDocument() to create new documents in production.
func newDocumentWithSections(filename DocumentFilename, sections SectionCollection) *Document {
	doc := Document{filename: filename}
	doc.sections = sections
	return &doc
}

// ScaffoldDocument provides new Documents for testing.
func ScaffoldDocument(data DocumentScaffold) *Document {
	if data.FileName == "" {
		data.FileName = "default.md"
	}
	if data.Content == "" {
		data.Content = "# Title\ndefault content"
	}
	return newDocumentWithText(DocumentFilename(data.FileName), data.Content)
}

// AllSections returns all the TikiSections that make up this document,
// including the title section.
func (doc *Document) AllSections() (result SectionCollection) {
	return doc.sections
}

// AppendSection provides a new Document with the given Section appended.
func (doc *Document) AppendSection(section *Section) *Document {
	// add an empty line to the last section
	lastSection := doc.sections[len(doc.sections)-1]
	newLastSection := lastSection.AppendText("\n")
	replacedSections := doc.sections.Replace(lastSection, &newLastSection)
	// add the new section
	newSections := append(replacedSections, section)
	return newDocumentWithSections(doc.filename, newSections)
}

// Content returns the content of this document.
func (doc *Document) Content() string {
	return doc.sections.Text()
}

// ContentSections provides the content sections of this document.
func (doc *Document) ContentSections() SectionCollection {
	return doc.sections[1:]
}

// FindSectionWithTitle provides the section with the given title,
// or nil in this document doesn't contain such a section.
func (doc *Document) FindSectionWithTitle(title string) (*Section, error) {
	return doc.sections.FindByTitle(title)
}

// FileName returns the file path (handle + extension) of this Document.
func (doc *Document) FileName() DocumentFilename {
	return doc.filename
}

// ID provides the unique ID of this document.
func (doc *Document) ID() string {
	filename := string(doc.filename)
	return strings.TrimSuffix(filename, path.Ext(filename))
}

// RemoveSection provides a copy of this Document that contains all its sections except the given one.
func (doc *Document) RemoveSection(section *Section) *Document {
	return newDocumentWithSections(doc.filename, doc.AllSections().Remove(section))
}

// ReplaceSection provides a new Document that is like this one
// and has the given old section replaced with the given new section.
func (doc *Document) ReplaceSection(oldSection *Section, newSection *Section) *Document {
	return newDocumentWithSections(doc.filename, doc.AllSections().Replace(oldSection, newSection))
}

// TikiLinks returns the TikiLinks in this Document.
func (doc *Document) TikiLinks(dc DocumentCollection) (result TikiLinkCollection, err error) {
	for i := range doc.sections {
		section := doc.sections[i]
		sectionTitle, err := section.Title()
		if err != nil {
			return result, fmt.Errorf("cannot determine the TikiLinks of document %q: %w", doc.filename, err)
		}
		if sectionTitle == "occurrences" {
			// links inside existing "occurrences" sections don't count
			continue
		}
		links, err := doc.sections[i].TikiLinks(dc)
		if err != nil {
			return result, fmt.Errorf("cannot determine the TikiLinks of document %q: %w", doc.filename, err)
		}
		result = append(result, links...)
	}
	return result, nil
}

// TitleSection provides the section before the content sections start.
func (doc *Document) TitleSection() *Section {
	return doc.sections[0]
}

// URL provides the URL of this Document within its TikiBase.
func (doc *Document) URL() string {
	return string(doc.filename)
}
