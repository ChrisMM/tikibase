package storage

// TikiLink represents a hyperlink to another TikiDocument
type TikiLink struct {

	// the section in which this link is located
	sourceSection TikiSection

	// the document that this link points to
	targetDocument TikiDocument

	// the link title
	title string
}

// NewTikiLink creates a new TikiLink instance.
func NewTikiLink(title string, sourceSection TikiSection, targetDocument TikiDocument) TikiLink {
	return TikiLink{title: title, sourceSection: sourceSection, targetDocument: targetDocument}
}

// SourceSection provides the TikiSection in which this link is located.
func (tl TikiLink) SourceSection() TikiSection {
	return tl.sourceSection
}

// TargetDocument provides the TikiDocument that this link points to.
func (tl TikiLink) TargetDocument() TikiDocument {
	return tl.targetDocument
}

// Title provides the human-readable title of this link.
func (tl TikiLink) Title() string {
	return tl.title
}
