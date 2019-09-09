package mentions

import (
	"fmt"

	"github.com/kevgo/tikibase/domain"
	"github.com/pkg/errors"
)

// Run executes the "mentions" command in the given directory.
func Run(dir string) error {
	tb, err := domain.NewDirectoryTikiBase(dir)
	if err != nil {
		return err
	}

	docs, err := tb.Documents()
	if err != nil {
		return errors.Wrapf(err, "cannot get documents of TikiBase")
	}
	fmt.Printf("%d documents found\n", len(docs))

	allLinks, err := tb.TikiLinks()
	if err != nil {
		return errors.Wrapf(err, "cannot get links of TikiBase")
	}
	fmt.Printf("%d total links found\n", len(allLinks))

	linksToDocs := LinksToDocs(allLinks)
	fmt.Printf("%d linked documents found", len(linksToDocs))

	for _, doc := range docs {
		fileName := doc.FileName()
		l := linksToDocs[fileName]
		fmt.Printf("- %s: %d references\n", fileName, len(l))
	}

	return nil
}
