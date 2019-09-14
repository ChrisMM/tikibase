package domain

// SectionCollection is a collection of Sections.
type SectionCollection []Section

// ScaffoldSectionCollection creates new SectionCollection instances for testing.
func ScaffoldSectionCollection(data []SectionScaffold) (result SectionCollection) {
	for _, sectionData := range data {
		result = append(result, ScaffoldSection(sectionData))
	}
	return result
}

// Replace provides a new SectionCollection where the given old section is replaced with the given new section.
func (sc SectionCollection) Replace(oldSection, newSection Section) (result SectionCollection) {
	for _, section := range sc {
		if section == oldSection {
			result = append(result, newSection)
		} else {
			result = append(result, section)
		}
	}
	return result
}

// Text provides the full text of this SectionCollection.
func (sc SectionCollection) Text() (result string) {
	for _, section := range sc {
		result += string(section.Content())
	}
	return result
}
