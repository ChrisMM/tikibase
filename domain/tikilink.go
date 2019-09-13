package domain

// TikiLink represents a hyperlink from one Document to another Document
type TikiLink struct {

	// the section in which this link is located
	sourceSection TikiSection

	// the document that this link points to
	targetDocument Document

	// the link title
	title string
}

// TikiLinkScaffold defines arguments for ScaffoldTikiLink.
type TikiLinkScaffold struct {
	SourceSection  TikiSection
	TargetDocument Document
	Title          string
}

// newTikiLink creates a new TikiLink instance.
func newTikiLink(title string, sourceSection TikiSection, targetDocument Document) TikiLink {
	return TikiLink{title: title, sourceSection: sourceSection, targetDocument: targetDocument}
}

// ScaffoldTikiLink provides TikiLink instances for testing.
func ScaffoldTikiLink(data TikiLinkScaffold) TikiLink {
	return newTikiLink(data.Title, data.SourceSection, data.TargetDocument)
}

// SourceSection provides the TikiSection in which this link is located.
func (tl TikiLink) SourceSection() TikiSection {
	return tl.sourceSection
}

// TargetDocument provides the Document that this link points to.
func (tl TikiLink) TargetDocument() Document {
	return tl.targetDocument
}

// Title provides the human-readable title of this link.
func (tl TikiLink) Title() string {
	return tl.title
}
