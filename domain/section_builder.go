package domain

import "strings"

// SectionBuilder builds TikiSections out of lines.
type SectionBuilder struct {
	lines []string
}

// NewSectionBuilder creates a new TikiSectionBuilder with the given title
func NewSectionBuilder(title string) SectionBuilder {
	return SectionBuilder{lines: []string{title}}
}

// AddLine adds a content line to this TikiSectionBuilder.
func (tsb *SectionBuilder) AddLine(line string) {
	tsb.lines = append(tsb.lines, line)
}

// Section provides the TikiSection that has been built up so far.
func (tsb *SectionBuilder) Section() Section {
	return Section{content: SectionContent(strings.Join(tsb.lines, "\n"))}
}
