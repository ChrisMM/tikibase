package occurrences

import (
	"fmt"
	"log"

	"github.com/kevgo/tikibase/domain"
	"github.com/pkg/errors"
)

// Run executes the "occurrences" command in the given directory.
func Run(dir string) error {
	tb, err := domain.NewTikiBase(dir)
	if err != nil {
		return err
	}

	docs, err := tb.Documents()
	if err != nil {
		return errors.Wrap(err, "cannot get documents of TikiBase")
	}

	allLinks, err := docs.TikiLinks()
	if err != nil {
		return errors.Wrap(err, "cannot get links of TikiBase")
	}

	linksToDocs := allLinks.GroupByTarget()

	for i := range docs {
		fileName := docs[i].FileName()
		linksToDoc := linksToDocs[fileName]
		oldOccurrencesSection, err := docs[i].FindSectionWithTitle("occurrences")
		if err != nil {
			return errors.Wrapf(err, "error finding existing occurrences sections in document '%s'", fileName)
		}
		newOccurrencesSection, err := RenderOccurrencesSection(linksToDoc, &docs[i])
		if err != nil {
			return errors.Wrapf(err, "error rendering new occurrences sections for document '%s'", fileName)
		}
		var doc2 domain.Document
		switch {
		case len(linksToDoc) == 0 && oldOccurrencesSection == nil:
			// no links to this doc and no existing occurrences section --> ignore this file
			continue
		case len(linksToDoc) == 0 && oldOccurrencesSection != nil:
			// no links to this doc but existing occurrences section --> delete the existing "occurrences" section
			doc2 = docs[i].RemoveSection(oldOccurrencesSection)
		case len(linksToDoc) > 0 && oldOccurrencesSection != nil:
			// links to this doc and existing "occurrences" section --> replace the existing "occurrences" section
			doc2 = docs[i].ReplaceSection(oldOccurrencesSection, newOccurrencesSection)
		case len(linksToDoc) > 0 && oldOccurrencesSection == nil:
			// links to this doc and no existing "occurrences" section --> append a new "occurrences" section
			doc2 = docs[i].AppendSection(newOccurrencesSection)
		}
		fmt.Print(".")
		err = tb.SaveDocument(doc2)
		if err != nil {
			log.Fatalf("cannot update document %s: %v", fileName, err)
		}
	}

	fmt.Printf("\n\nprocessed %d TikiLinks in %d documents...\n", len(allLinks), len(docs))
	return nil
}
