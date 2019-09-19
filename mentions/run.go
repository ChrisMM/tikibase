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
	fmt.Printf("%d documents found\n", len(docs))

	allLinks, err := docs.TikiLinks()
	if err != nil {
		return errors.Wrapf(err, "cannot get links of TikiBase")
	}
	fmt.Printf("%d total links found\n", len(allLinks))

	linksToDocs := allLinks.GroupByTarget()
	fmt.Printf("%d linked documents found\n", len(linksToDocs))

	for i := range docs {
		fileName := docs[i].FileName()
		linksToDoc := linksToDocs[fileName]
		fmt.Printf("- %s: %d references\n", fileName, len(linksToDoc))
		if len(linksToDoc) == 0 {
			continue
		}
		mentionsSection := RenderMentionsSection(linksToDoc, &docs[i])
		doc2 := docs[i].AppendSection(mentionsSection)
		err := tb.SaveDocument(doc2)
		if err != nil {
			log.Fatalf("cannot update document %s: %v", fileName, err)
		}
	}

	return nil
}
