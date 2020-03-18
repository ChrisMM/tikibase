package find

import (
	"fmt"
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
	for i := range docs {
		section, err := docs[i].FindSectionWithTitle("what is it")
		if err != nil || section == nil {
			continue
		}
		sectionContent := section.Content()
		for i := range types {
			if !strings.Contains(sectionContent, types[i]) {
				continue
			}
		}
		title, err := docs[i].TitleSection().Title()
		if err != nil {
			return result, fmt.Errorf("error getting the title of a document, please run \"tikibase check\" to investigate")
		}
		result = append(result, title)
	}
	return result, nil
}
