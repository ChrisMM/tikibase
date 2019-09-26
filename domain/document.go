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
func (d *Document) AllSections() (result *SectionCollection) {
	return &d.sections
}

// AppendSection provides a new Document with the given Section appended.
func (d *Document) AppendSection(section Section) *Document {
	// add an empty line to the last section
	lastSection := d.sections[len(d.sections)-1]
	newLastSection := lastSection.AppendLine("\n")
	replacedSections := d.sections.Replace(&lastSection, newLastSection)

	// add the new section
	newSections := append(replacedSections, section)
	return newDocument(d.filename, newSections.Text())
}

// Content returns the content of this document.
func (d *Document) Content() string {
	return d.sections.Text()
}

// ContentSections provides the content sections of this document.
func (d *Document) ContentSections() *SectionCollection {
	result := d.sections[1:]
	return &result
}

// FindSectionWithTitle provides the section with the given title,
// or nil in this document doesn't contain such a section.
func (d *Document) FindSectionWithTitle(title string) (*Section, error) {
	return d.sections.FindByTitle(title)
}

// FileName returns the file path (handle + extension) of this Document.
func (d *Document) FileName() DocumentFilename {
	return d.filename
}

// RemoveSection provides a copy of this Document that contains all its sections except the given one.
func (d *Document) RemoveSection(section *Section) *Document {
	newSections := d.AllSections().Remove(section)
	return newDocument(d.filename, newSections.Text())
}

// ReplaceSection provides a new Document that is like this one
// and has the given old section replaced with the given new section.
func (d *Document) ReplaceSection(oldSection *Section, newSection Section) *Document {
	newSections := d.AllSections().Replace(oldSection, newSection)
	return newDocument(d.filename, newSections.Text())
}

// TikiLinks returns the TikiLinks in this Document.
func (d *Document) TikiLinks(tdc DocumentCollection) (result TikiLinkCollection, err error) {
	for i := range d.sections {
		section := d.sections[i]
		sectionTitle, err := section.Title()
		if err != nil {
			return result, errors.Wrapf(err, "Cannot determine the TikiLinks of document '%s'", d.filename)
		}
		if sectionTitle == "occurrences" {
			// links inside existing "occurrences" sections don't count
			continue
		}
		links, err := d.sections[i].TikiLinks(tdc)
		if err != nil {
			return result, errors.Wrapf(err, "cannot determine the TikiLinks of document '%s'", d.filename)
		}
		result = append(result, links...)
	}
	return result, nil
}

// TitleSection provides the section before the content sections start.
func (d *Document) TitleSection() *Section {
	return &d.sections[0]
}

// URL provides the URL of this Document within its TikiBase.
func (d *Document) URL() string {
	return string(d.filename)
}
