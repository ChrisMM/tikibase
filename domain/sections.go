package domain

import (
	"fmt"
	"strings"

	"github.com/kevgo/tikibase/helpers"
)

// Sections is a collection of Sections.
type Sections []*Section

func newSectionCollection(content string, doc *Document) (result Sections) {
	sb := NewSectionBuilder("", doc)
	lines := helpers.CutStringIntoLines(content)
	for i := range lines {
		if strings.HasPrefix(lines[i], "#") {
			if sb.Len() > 0 {
				result = append(result, sb.Section())
			}
			sb = NewSectionBuilder(lines[i], doc)
		} else {
			sb.AddLine(lines[i])
		}
	}
	result = append(result, sb.Section())
	return result
}

// ScaffoldSections creates new Sections instances for testing.
func ScaffoldSections(data []SectionScaffold) (result Sections) {
	for i := range data {
		result = append(result, ScaffoldSection(data[i]))
	}
	return result
}

// FindByTitle provides the section with the given title
// or nil if none was found
func (sections Sections) FindByTitle(title string) (*Section, error) {
	for i := range sections {
		section := sections[i]
		sectionTitle, err := section.Title()
		if err != nil {
			return nil, fmt.Errorf("cannot find section with title %q: %w", title, err)
		}
		if sectionTitle == title {
			return section, nil
		}
	}
	return nil, nil
}

// Remove provides a copy of this SectionCollection with the given sections removed.
func (sections Sections) Remove(section *Section) (result Sections) {
	for i := range sections {
		if sections[i] != section {
			result = append(result, sections[i])
		}
	}
	return result
}

// Replace provides a new SectionCollection where the given old section is replaced with the given new section.
func (sections Sections) Replace(oldSection *Section, newSection *Section) (result Sections) {
	for i := range sections {
		if sections[i] == oldSection {
			result = append(result, newSection)
		} else {
			result = append(result, sections[i])
		}
	}
	return result
}

// Text provides the full text of this SectionCollection.
func (sections Sections) Text() (result string) {
	for i := range sections {
		result += sections[i].Content()
	}
	return result
}
