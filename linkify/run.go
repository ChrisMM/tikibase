package linkify

import (
	"fmt"
	"strings"

	"github.com/kevgo/tikibase/domain"
)

// Run creates missing links in this TikiBase.
func Run(dir string, log bool) (err error) {
	tikibase, err := domain.NewTikiBase(dir)
	if err != nil {
		return err
	}
	docs, err := tikibase.Documents()
	if err != nil {
		return fmt.Errorf("cannot get documents of TikiBase: %w", err)
	}

	// find all document titles
	titles := make(map[string]string) // text -> filename
	for d := range docs {
		names, err := docs[d].Names()
		if err != nil {
			return err
		}
		for n := range names {
			titles[names[n]] = docs[d].FileName()
		}
	}

	// scan all documents for missing link titles
	for d := range docs {
		docContent := docs[d].Content()
		if log {
			fmt.Print(".")
		}
		docTitle, err := docs[d].Title()
		if err != nil {
			return err
		}
		linkified := docContent
		for title := range titles {
			if strings.EqualFold(title, docTitle) {
				continue
			}
			linkified = linkifyDoc(linkified, title, titles[title])
		}
		if linkified != docContent {
			err = tikibase.UpdateDocument(docs[d], linkified)
			if err != nil {
				return err
			}
		}
	}
	if log {
		fmt.Println()
	}
	return nil
}
