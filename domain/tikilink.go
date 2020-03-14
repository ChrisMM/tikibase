package domain

// TikiLink represents a hyperlink from one Document to another Document
type TikiLink struct {
	Link

	// the document that this link points to
	targetDocument *Document
}

// TikiLinkScaffold defines arguments for ScaffoldTikiLink.
type TikiLinkScaffold struct {
	SourceSection  *Section
	TargetDocument *Document
	Title          string
}

// newTikiLink creates a new TikiLink instance.
func newTikiLink(title string, sourceSection *Section, targetDocument *Document) *TikiLink {
	result := &TikiLink{targetDocument: targetDocument}
	result.title = title
	result.sourceSection = sourceSection
	return result
}

// ScaffoldTikiLink provides TikiLink instances for testing.
func ScaffoldTikiLink(data TikiLinkScaffold) *TikiLink {
	if data.Title == "" {
		data.Title = "default title"
	}
	return newTikiLink(data.Title, data.SourceSection, data.TargetDocument)
}

// TargetDocument provides the Document that this link points to.
func (link *TikiLink) TargetDocument() *Document {
	return link.targetDocument
}
