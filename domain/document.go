package domain

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

// Document represents a MarkDown file in the document base.
// Create new instances via DocumentCollection.CreateDocument
type Document struct {

	// fileame is the unique identifier for this document.
	filename string

	// the textual content of the document
	sections Sections
}

// DocumentScaffold is for easy scaffolding of Documents in tests.
// Don't use this in production code.
type DocumentScaffold struct {
	FileName string
	Content  string
}

// newDocumentWithText creates a new Document instance with the given textual content.
// This constructor is internal to this module,
// call (TikiBase).CreateDocument() to create new documents in production.
func newDocumentWithText(filename string, content string) *Document {
	doc := Document{filename: filename}
	doc.sections = newSectionCollection(content, &doc)
	return &doc
}

// newDocumentWithSections creates a new Document instance with the given pre-parsed sections.
// This constructor is internal to this module,
// call (TikiBase).CreateDocument() to create new documents in production.
func newDocumentWithSections(filename string, sections Sections) *Document {
	doc := Document{filename, sections}
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
	return newDocumentWithText(data.FileName, data.Content)
}

// AllSections returns all the TikiSections that make up this document,
// including the title section.
func (doc *Document) AllSections() (result Sections) {
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
func (doc *Document) ContentSections() Sections {
	return doc.sections[1:]
}

// FindSectionWithTitle provides the section with the given title,
// or nil in this document doesn't contain such a section.
func (doc *Document) FindSectionWithTitle(title string) (*Section, error) {
	return doc.sections.FindByTitle(title)
}

// FileName returns the file path (handle + extension) of this Document.
func (doc *Document) FileName() string {
	return doc.filename
}

// ID provides the unique ID of this document.
func (doc *Document) ID() string {
	return strings.TrimSuffix(doc.filename, filepath.Ext(doc.filename))
}

// LinkedDocs provides the documents that this document links to.
func (doc *Document) LinkedDocs(docs Documents) (result []*Document, err error) {
	links, err := doc.TikiLinks(docs)
	if err != nil {
		return result, fmt.Errorf("cannot determine the linked docs of %q: %w", doc.FileName(), err)
	}
	unifier := make(map[string]*Document)
	for l := range links {
		target := links[l].TargetDocument()
		unifier[target.FileName()] = target
	}
	for filename := range unifier {
		result = append(result, unifier[filename])
	}
	return result, nil
}

// Links provides all Links in this document.
func (doc *Document) Links() (result []Link) {
	for i := range doc.sections {
		result = append(result, doc.sections[i].Links()...)
	}
	return result
}

// Names provides the human-readable names of this document.
func (doc *Document) Names() (result []string, err error) {
	title, err := doc.Title()
	if err != nil {
		return result, fmt.Errorf("cannot determine names: %w", err)
	}
	namesTitleOnce.Do(func() { namesTitleRE = regexp.MustCompile(`\(.*\)$`) })
	abbreviations := namesTitleRE.FindStringSubmatch(title)
	if len(abbreviations) > 0 {
		result = []string{
			strings.TrimSpace(strings.Replace(title, abbreviations[0], "", 1)),
			strings.Replace(strings.Replace(abbreviations[0], "(", "", 1), ")", "", 1),
		}
	} else {
		result = []string{title}
	}
	return result, nil
}

//nolint:gochecknoglobals
var namesTitleOnce sync.Once

//nolint:gochecknoglobals
var namesTitleRE *regexp.Regexp

// RemoveSection provides a copy of this Document that contains all its sections except the given one.
func (doc *Document) RemoveSection(section *Section) *Document {
	return newDocumentWithSections(doc.filename, doc.AllSections().Remove(section))
}

// ReplaceSection provides a new Document that is like this one
// and has the given old section replaced with the given new section.
func (doc *Document) ReplaceSection(oldSection *Section, newSection *Section) *Document {
	return newDocumentWithSections(doc.filename, doc.AllSections().Replace(oldSection, newSection))
}

// Title provides the title of this document.
func (doc *Document) Title() (result string, err error) {
	return doc.TitleSection().Title()
}

// TikiLinks returns the TikiLinks in this Document.
func (doc *Document) TikiLinks(dc Documents) (result TikiLinks, err error) {
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
	return doc.filename
}
