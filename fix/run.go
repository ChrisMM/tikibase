package fix

import (
	"fmt"

	"github.com/kevgo/tikibase/domain"
	"github.com/kevgo/tikibase/occurrences"
)

// Run executes the fix" command in the given directory.
//nolint:funlen
func Run(dir string) (err error) {
	tikibase, err := domain.NewTikiBase(dir)
	if err != nil {
		return err
	}
	docs, err := tikibase.Documents()
	if err != nil {
		return err
	}
	allLinks, err := docs.TikiLinks()
	if err != nil {
		return err
	}
	linksToDocs := allLinks.GroupByTarget()
	for i := range docs {
		fileName := docs[i].FileName()
		linksInDoc, err := docs[i].TikiLinks(docs)
		if err != nil {
			return fmt.Errorf("cannot determine occurrences for document %q", err)
		}
		allOccurrencesLinks := linksToDocs[fileName]
		missingOccurrencesLinks := allOccurrencesLinks.RemoveLinksFromDocs(linksInDoc.ReferencedDocs())
		dedupedOccurrencesLinks := missingOccurrencesLinks.CombineLinksFromSameDocuments()
		dedupedOccurrencesLinks.SortBySourceDocumentTitle()
		newOccurrencesSection, err := occurrences.CreateSection(dedupedOccurrencesLinks, docs[i])
		if err != nil {
			return fmt.Errorf("error creating new occurrences sections for document %q: %w", fileName, err)
		}
		existingOccurrencesSection, err := docs[i].FindSectionWithTitle("occurrences")
		if err != nil {
			return fmt.Errorf("error finding existing occurrences sections in document %q: %w", fileName, err)
		}
		var newDoc *domain.Document
		switch {
		case len(dedupedOccurrencesLinks) == 0 && existingOccurrencesSection == nil:
			continue
		case len(dedupedOccurrencesLinks) == 0 && existingOccurrencesSection != nil:
			newDoc = docs[i].RemoveSection(existingOccurrencesSection)
		case len(dedupedOccurrencesLinks) > 0 && existingOccurrencesSection != nil:
			if newOccurrencesSection.Content() == existingOccurrencesSection.Content() {
				continue
			}
			newDoc = docs[i].ReplaceSection(existingOccurrencesSection, newOccurrencesSection)
		case len(dedupedOccurrencesLinks) > 0 && existingOccurrencesSection == nil:
			newDoc = docs[i].AppendSection(newOccurrencesSection)
		}
		err = tikibase.SaveDocument(newDoc)
		if err != nil {
			return fmt.Errorf("cannot update document %s: %w", fileName, err)
		}
	}

	return nil
}
