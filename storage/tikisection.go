package storage

// TikiSection represents a section in a TikiDocument.
type TikiSection struct {
	// the textual content of this TikiSection
	// including the title line
	content string
}

// NewTikiSection creates a new TikiSection with the given content.
func NewTikiSection(content string) TikiSection {
	return TikiSection{content: content}
}

// Content returns the content of the entire section as a block.
func (ts TikiSection) Content() string {
	return ts.content
}
