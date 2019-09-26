package domain

import "strings"

// SectionBuilder builds TikiSections out of lines.
// Lines must be provided as-is, i.e. contain the ending newlines.
type SectionBuilder struct {
	content  *strings.Builder
	document *Document
}

// NewSectionBuilder creates a new TikiSectionBuilder with the given title
func NewSectionBuilder(title string, doc *Document) SectionBuilder {
	builder := SectionBuilder{content: &strings.Builder{}, document: doc}
	builder.content.WriteString(title)
	return builder
}

// AddLine adds a content line to this TikiSectionBuilder.
func (sb *SectionBuilder) AddLine(line string) {
	sb.content.WriteString(line)
}

// Len provides the length of the section content accumulated so far.
func (sb *SectionBuilder) Len() int {
	return sb.content.Len()
}

// Section provides the TikiSection that has been built up so far.
func (sb *SectionBuilder) Section() *Section {
	return &Section{content: SectionContent(sb.content.String()), document: sb.document}
}
