package domain

// TikiLink represents a hyperlink from one Document to another Document
type TikiLink struct {

	// the section in which this link is located
	sourceSection *Section

	// the document that this link points to
	targetDocument *Document

	// the link title
	title string
}

// TikiLinkScaffold defines arguments for ScaffoldTikiLink.
type TikiLinkScaffold struct {
	SourceSection  *Section
	TargetDocument *Document
	Title          string
}

// newTikiLink creates a new TikiLink instance.
func newTikiLink(title string, sourceSection *Section, targetDocument *Document) *TikiLink {
	return &TikiLink{title: title, sourceSection: sourceSection, targetDocument: targetDocument}
}

// ScaffoldTikiLink provides TikiLink instances for testing.
func ScaffoldTikiLink(data TikiLinkScaffold) *TikiLink {
	if data.Title == "" {
		data.Title = "default title"
	}
	return newTikiLink(data.Title, data.SourceSection, data.TargetDocument)
}

// SourceSection provides the TikiSection in which this link is located.
func (link *TikiLink) SourceSection() *Section {
	return link.sourceSection
}

// TargetDocument provides the Document that this link points to.
func (link *TikiLink) TargetDocument() *Document {
	return link.targetDocument
}

// Title provides the human-readable title of this link.
func (link *TikiLink) Title() string {
	return link.title
}
