package remove

import (
	"fmt"

	"github.com/kevgo/tikibase/domain"
)

// Run performs the "delete" command.
func Run(dir, filename string) error {
	tikibase, err := domain.NewTikiBase(dir)
	if err != nil {
		return err
	}
	docs, err := tikibase.Documents()
	if err != nil {
		return err
	}
	for d := range docs {
		content := docs[d].Content()
		newContent := removeLinksToFile(filename, content)
		if newContent != content {
			err := tikibase.UpdateDocument(docs[d], newContent)
			if err != nil {
				return fmt.Errorf("cannot update document %q: %w", docs[d].FileName(), err)
			}
		}
	}
	return nil
}
