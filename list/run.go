package list

import (
	"fmt"
	"os"

	"github.com/kevgo/tikibase/domain"
	"github.com/kevgo/tikibase/helpers"
)

// Run performs the search.
func Run(dir string, types []string) (result []string, err error) {
	tikibase, err := domain.NewTikiBase(dir)
	if err != nil {
		return result, err
	}
	docs, err := tikibase.Documents()
	if err != nil {
		return result, fmt.Errorf("cannot get documents of TikiBase: %w", err)
	}
	i := 0
	for _, doc := range docs {
		section, err := doc.FindSectionWithTitle("what is it")
		if err != nil {
			continue
		}
		if section == nil {
			continue
		}
		sectionContent := string(section.Content())
		if !helpers.ContainsStrings(sectionContent, types) {
			continue
		}
		i++
		title, err := doc.TitleSection().Title()
		if err != nil {
			fmt.Printf("Error getting the title of document %s: %v\n", doc.ID(), err)
			os.Exit(1)
		}
		result = append(result, title)
	}
	return result, nil
}
