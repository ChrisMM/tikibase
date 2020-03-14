package domain

// Link represents a hyperlink inside a document.
type Link struct {
	// the section in which this link is located
	sourceSection *Section

	// the link target
	target string

	// the link title
	title string
}

// SourceSection provides the TikiSection in which this link is located.
func (link *Link) SourceSection() *Section {
	return link.sourceSection
}

// Target provides the link target.
func (link *Link) Target() string {
	return link.target
}

// Title provides the human-readable title of this link.
func (link *Link) Title() string {
	return link.title
}

// ScaffoldLink provides Link instances for testing.
func ScaffoldLink(target string) Link {
	return Link{target: target, title: "scaffolded link"}
}
