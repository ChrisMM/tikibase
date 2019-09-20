package mentions

import (
	"fmt"
	"log"

	"github.com/kevgo/tikibase/domain"
	"github.com/pkg/errors"
)

// Run executes the "mentions" command in the given directory.
func Run(dir string) error {
	tb, err := domain.NewTikiBase(dir)
	if err != nil {
		return err
	}

	docs, err := tb.Documents()
	if err != nil {
		return errors.Wrapf(err, "cannot get documents of TikiBase")
	}

	allLinks, err := docs.TikiLinks()
	if err != nil {
		return errors.Wrapf(err, "cannot get links of TikiBase")
	}
	fmt.Printf("processing %d links in %d documents...\n", len(allLinks), len(docs))

	linksToDocs := allLinks.GroupByTarget()

	for i := range docs {
		fileName := docs[i].FileName()
		linksToDoc := linksToDocs[fileName]
		fmt.Printf("- %s: %d references\n", fileName, len(linksToDoc))
		oldMentionsSection := docs[i].FindSectionWithTitle("mentions")
		newMentionsSection := RenderMentionsSection(linksToDoc, &docs[i])
		var doc2 domain.Document
		switch {
		case len(linksToDoc) == 0 && oldMentionsSection == nil:
			// no links to this doc and no existing mentions section --> ignore this file
			continue
		case len(linksToDoc) == 0 && oldMentionsSection != nil:
			// no links to this doc but existing mentions section --> delete the existing "mentions" section
			doc2 = docs[i].RemoveSection(oldMentionsSection)
		case len(linksToDoc) > 0 && oldMentionsSection != nil:
			// links to this doc and existing "mentions" section --> replace the existing "mentions" section
			doc2 = docs[i].ReplaceSection(oldMentionsSection, newMentionsSection)
		case len(linksToDoc) > 0 && oldMentionsSection == nil:
			// links to this doc and no existing "mentions" section --> append a new "mentions" section
			doc2 = docs[i].AppendSection(newMentionsSection)
		}
		err := tb.SaveDocument(doc2)
		if err != nil {
			log.Fatalf("cannot update document %s: %v", fileName, err)
		}
	}

	return nil
}
