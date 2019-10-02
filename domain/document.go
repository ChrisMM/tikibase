package domain

import "github.com/pkg/errors"

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

// newDocument creates a new Document instance.
// This constructor is internal to this module,
// call (TikiBase).CreateDocument() to create new documents in production.
func newDocument(filename DocumentFilename, content string) *Document {
	doc := Document{filename: filename}
	doc.sections = newSectionCollection(content, &doc)
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
	return newDocument(DocumentFilename(data.FileName), data.Content)
}

// AllSections returns all the TikiSections that make up this document,
// including the title section.
func (doc *Document) AllSections() (result *SectionCollection) {
	return &doc.sections
}

// AppendSection provides a new Document with the given Section appended.
func (doc *Document) AppendSection(section *Section) *Document {
	// add an empty line to the last section
	lastSection := doc.sections[len(doc.sections)-1]
	newLastSection := lastSection.AppendLine("\n")
	replacedSections := doc.sections.Replace(lastSection, &newLastSection)

	// add the new section
	newSections := append(replacedSections, section)
	return newDocument(doc.filename, newSections.Text())
}

// Content returns the content of this document.
func (doc *Document) Content() string {
	return doc.sections.Text()
}

// ContentSections provides the content sections of this document.
func (doc *Document) ContentSections() *SectionCollection {
	result := doc.sections[1:]
	return &result
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

// RemoveSection provides a copy of this Document that contains all its sections except the given one.
func (doc *Document) RemoveSection(section *Section) *Document {
	newSections := doc.AllSections().Remove(section)
	return newDocument(doc.filename, newSections.Text())
}

// ReplaceSection provides a new Document that is like this one
// and has the given old section replaced with the given new section.
func (doc *Document) ReplaceSection(oldSection *Section, newSection *Section) *Document {
	newSections := doc.AllSections().Replace(oldSection, newSection)
	return newDocument(doc.filename, newSections.Text())
}

// TikiLinks returns the TikiLinks in this Document.
func (doc *Document) TikiLinks(tdc DocumentCollection) (result TikiLinkCollection, err error) {
	for i := range doc.sections {
		section := doc.sections[i]
		sectionTitle, err := section.Title()
		if err != nil {
			return result, errors.Wrapf(err, "Cannot determine the TikiLinks of document %q", doc.filename)
		}
		if sectionTitle == "occurrences" {
			// links inside existing "occurrences" sections don't count
			continue
		}
		links, err := doc.sections[i].TikiLinks(tdc)
		if err != nil {
			return result, errors.Wrapf(err, "cannot determine the TikiLinks of document %q", doc.filename)
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
