package domain

import (
	"strings"

	"github.com/kevgo/tikibase/helpers"
)

// SectionCollection is a collection of Sections.
type SectionCollection []Section

func newSectionCollection(content string, doc *Document) (result SectionCollection) {
	tsb := NewSectionBuilder("", doc)
	lines := helpers.CutStringIntoLines(content)
	for i := range lines {
		if strings.HasPrefix(lines[i], "#") {
			if tsb.Len() > 0 {
				result = append(result, tsb.Section())
			}
			tsb = NewSectionBuilder(lines[i], doc)
		} else {
			tsb.AddLine(lines[i])
		}
	}
	result = append(result, tsb.Section())
	return result
}

// ScaffoldSectionCollection creates new SectionCollection instances for testing.
func ScaffoldSectionCollection(data []SectionScaffold) (result SectionCollection) {
	for i := range data {
		result = append(result, ScaffoldSection(data[i]))
	}
	return result
}

// Equal indicates whether this SectionCollection has the same content as the given one.
func (sc SectionCollection) Equal(other SectionCollection) bool {
	if len(sc) != len(other) {
		return false
	}
	for i := range sc {
		if sc[i] != other[i] {
			return false
		}
	}

	return true
}

// Replace provides a new SectionCollection where the given old section is replaced with the given new section.
func (sc SectionCollection) Replace(oldSection, newSection Section) (result SectionCollection) {
	for i := range sc {
		if sc[i] == oldSection {
			result = append(result, newSection)
		} else {
			result = append(result, sc[i])
		}
	}
	return result
}

// Text provides the full text of this SectionCollection.
func (sc SectionCollection) Text() (result string) {
	for i := range sc {
		result += string(sc[i].Content())
	}
	return result
}
