package linkify

import (
	"fmt"

	"github.com/kevgo/tikibase/domain"
)

type logger func()

// Run creates missing links in this TikiBase.
func Run(dir string, log logger) (err error) {
	tikibase, err := domain.NewTikiBase(dir)
	if err != nil {
		return err
	}
	docs, err := tikibase.Documents()
	if err != nil {
		return fmt.Errorf("cannot get documents of TikiBase: %w", err)
	}

	// find all document names
	mappings, err := docsMappings(docs)
	if err != nil {
		return err
	}

	// scan all documents for missing link titles
	for d := range docs {
		linkified := linkifyDoc(docs[d], mappings)
		if linkified != docs[d].Content() {
			err = tikibase.UpdateDocument(docs[d], linkified)
			if err != nil {
				return err
			}
		}
		log()
	}
	return nil
}
