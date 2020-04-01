package linkify

import (
	"fmt"
	"runtime"

	"github.com/kevgo/tikibase/domain"
	"golang.org/x/sync/errgroup"
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

	threadCount := runtime.NumCPU() * 4 // 4x concurrency per core seems a good compromise given the workers do file system operations
	var group errgroup.Group
	docsChan := make(chan *domain.Document)
	for i := 0; i < threadCount; i++ {
		group.Go(func() error {
			for doc := range docsChan {
				linkified, err := linkifyDoc(doc, docs, mappings)
				if err != nil {
					return fmt.Errorf("cannot linkify document %q: %w", doc.FileName(), err)
				}
				if linkified != doc.Content() {
					err = tikibase.UpdateDocument(doc, linkified)
					if err != nil {
						return fmt.Errorf("cannot update document %q: %w", doc.FileName(), err)
					}
				}
			}
			return nil
		})
	}
	for d := range docs {
		docsChan <- docs[d]
	}
	close(docsChan)

	return group.Wait()
}
