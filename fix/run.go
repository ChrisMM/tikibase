package fix

import (
	"fmt"

	"github.com/kevgo/tikibase/domain"
)

// Run executes the fix" command in the given directory.
//nolint:funlen
func Run(dir string) (docsCount, createdCount, updatedCount, deletedCount int, err error) {
	tikibase, err := domain.NewTikiBase(dir)
	if err != nil {
		return 0, 0, 0, 0, err
	}
	docs, err := tikibase.Documents()
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("cannot get documents of TikiBase: %w", err)
	}
	allLinks, err := docs.TikiLinks()
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("cannot get links of TikiBase: %w", err)
	}
	linksToDocs := allLinks.GroupByTarget()
	for i := range docs {
		docsCount++
		fileName := docs[i].FileName()
		linksInDoc, err := docs[i].TikiLinks(docs)
		if err != nil {
			return 0, 0, 0, 0, fmt.Errorf("cannot determine occurrences for document %q", err)
		}
		allOccurrencesLinks := linksToDocs[fileName]
		missingOccurrencesLinks := allOccurrencesLinks.RemoveLinksFromDocs(linksInDoc.ReferencedDocs())
		dedupedOccurrencesLinks := missingOccurrencesLinks.CombineLinksFromSameDocuments()
		dedupedOccurrencesLinks.SortBySourceDocumentTitle()
		newOccurrencesSection, err := CreateOccurrencesSection(dedupedOccurrencesLinks, docs[i])
		if err != nil {
			return 0, 0, 0, 0, fmt.Errorf("error creating new occurrences sections for document %q: %w", fileName, err)
		}
		existingOccurrencesSection, err := docs[i].FindSectionWithTitle("occurrences")
		if err != nil {
			return 0, 0, 0, 0, fmt.Errorf("error finding existing occurrences sections in document %q: %w", fileName, err)
		}
		var newDoc *domain.Document
		switch {
		case len(dedupedOccurrencesLinks) == 0 && existingOccurrencesSection == nil:
			continue
		case len(dedupedOccurrencesLinks) == 0 && existingOccurrencesSection != nil:
			deletedCount++
			newDoc = docs[i].RemoveSection(existingOccurrencesSection)
		case len(dedupedOccurrencesLinks) > 0 && existingOccurrencesSection != nil:
			if newOccurrencesSection.Content() == existingOccurrencesSection.Content() {
				continue
			}
			updatedCount++
			newDoc = docs[i].ReplaceSection(existingOccurrencesSection, newOccurrencesSection)
		case len(dedupedOccurrencesLinks) > 0 && existingOccurrencesSection == nil:
			createdCount++
			newDoc = docs[i].AppendSection(newOccurrencesSection)
		}
		err = tikibase.SaveDocument(newDoc)
		if err != nil {
			return docsCount, createdCount, updatedCount, deletedCount, fmt.Errorf("cannot update document %s: %w", fileName, err)
		}
	}

	return docsCount, createdCount, updatedCount, deletedCount, nil
}
