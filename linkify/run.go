package linkify

import (
	"fmt"
	"strings"

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

	// find all document titles
	docsNames := make(map[string]string) // doc name -> filename
	for d := range docs {
		docNames, err := docs[d].Names()
		if err != nil {
			return err
		}
		for dn := range docNames {
			docsNames[docNames[dn]] = docs[d].FileName()
		}
	}

	// scan all documents for missing link titles
	for d := range docs {
		docContent := docs[d].Content()
		docTitle, err := docs[d].Title()
		if err != nil {
			return err
		}
		linkified := docContent
		for title := range docsNames {
			if strings.EqualFold(title, docTitle) {
				continue
			}
			linkified = linkifyDoc(linkified, title, docsNames[title])
		}
		if linkified != docContent {
			err = tikibase.UpdateDocument(docs[d], linkified)
			if err != nil {
				return err
			}
		}
		log()
	}
	return nil
}
