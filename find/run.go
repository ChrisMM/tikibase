package find

import (
	"fmt"
	"os"
	"strings"

	"github.com/kevgo/tikibase/domain"
)

// Run performs the "find" command.
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
		if err != nil || section == nil {
			continue
		}
		sectionContent := string(section.Content())
		for i := range types {
			if !strings.Contains(sectionContent, types[i]) {
				continue
			}
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
