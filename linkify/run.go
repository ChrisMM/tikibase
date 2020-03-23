package linkify

import (
	"fmt"
	"strings"

	"github.com/kevgo/tikibase/domain"
)

// Run creates missing links in this TikiBase.
func Run(dir string) (err error) {
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
		linkified := docContent
		for title := range titles {
			docTitle, err := docs[d].TitleSection().Title()
			if err != nil {
				return err
			}
			if strings.EqualFold(title, docTitle) {
				continue
			}
			linkified = Linkify(linkified, title, titles[title])
		}
		if linkified != docContent {
			err = tikibase.UpdateDocument(docs[d], linkified)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
