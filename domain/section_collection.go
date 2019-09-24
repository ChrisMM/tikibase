package domain

import (
	"strings"

	"github.com/kevgo/tikibase/helpers"
	"github.com/pkg/errors"
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

// FindByTitle provides the section with the given title
// or nil if none was found
func (sc SectionCollection) FindByTitle(title string) (*Section, error) {
	for i := range sc {
		section := sc[i]
		sectionTitle, err := section.Title()
		if err != nil {
			return nil, errors.Wrapf(err, "cannot find section with title '%s'", title)
		}
		if sectionTitle == title {
			return &section, nil
		}
	}
	return nil, nil
}

// Remove provides a copy of this SectionCollection that contains all its sections except the given one.
func (sc SectionCollection) Remove(section *Section) (result SectionCollection) {
	for i := range sc {
		if sc[i] != *section {
			result = append(result, sc[i])
		}
	}
	return result
}

// Replace provides a new SectionCollection where the given old section is replaced with the given new section.
func (sc SectionCollection) Replace(oldSection *Section, newSection Section) (result SectionCollection) {
	for i := range sc {
		if sc[i] == *oldSection {
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
