package stats

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kevgo/tikibase/domain"
)

// Result contains statistics about a TikiBase.
type Result struct {
	DocsCount      int
	SectionsCount  int
	LinksCount     int
	ResourcesCount int
	SectionTypes   []string
}

// Run executes the "stats" command.
func Run(dir string) (result Result, err error) {
	tikibase, err := domain.NewTikiBase(dir)
	if err != nil {
		return
	}
	docFiles, resourceFiles, err := tikibase.Files()
	if err != nil {
		return
	}
	result.DocsCount = len(docFiles.FileNames())
	result.ResourcesCount = len(resourceFiles.FileNames())
	docs, err := docFiles.Documents()
	if err != nil {
		return
	}
	sectionTypes := make(map[string]struct{})
	for d := range docs {
		sections := docs[d].ContentSections()
		result.SectionsCount += len(sections)
		for s := range sections {
			result.LinksCount += len(sections[s].Links())
			title, err := sections[s].Title()
			if err != nil {
				return result, fmt.Errorf("please run \"tikibase check\" to investigate: %w", err)
			}
			sectionTypes[title] = struct{}{}
		}
	}
	for sectionType := range sectionTypes {
		result.SectionTypes = append(result.SectionTypes, sectionType)
	}
	sort.Slice(result.SectionTypes, func(i, j int) bool {
		return strings.ToLower(result.SectionTypes[i]) < strings.ToLower(result.SectionTypes[j])
	})
	return result, nil
}
