package storage

import "strings"

// TikiSectionBuilder builds TikiSections out of lines.
type TikiSectionBuilder struct {
	lines []string
}

// NewTikiSectionBuilder creates a new TikiSectionBuilder with the given title
func NewTikiSectionBuilder(title string) TikiSectionBuilder {
	return TikiSectionBuilder{lines: []string{title}}
}

// AddLine adds a content line to this TikiSectionBuilder.
func (tsb *TikiSectionBuilder) AddLine(line string) {
	tsb.lines = append(tsb.lines, line)
}

// Section provides the TikiSection that has been built up so far.
func (tsb *TikiSectionBuilder) Section() TikiSection {
	return TikiSection{content: strings.Join(tsb.lines, "\n") + "\n"}
}
