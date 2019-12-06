package occurrences

import (
	"fmt"

	"github.com/kevgo/tikibase/domain"
)

// Run executes the "occurrences" command in the given directory.
//nolint:funlen
func Run(dir string) error {
	output := NewDotOutput()
	tikibase, err := domain.NewTikiBase(dir)
	if err != nil {
		return err
	}
	docs, err := tikibase.Documents()
	if err != nil {
		return fmt.Errorf("cannot get documents of TikiBase: %w", err)
	}
	allLinks, err := docs.TikiLinks()
	if err != nil {
		return fmt.Errorf("cannot get links of TikiBase: %w", err)
	}
	linksToDocs := allLinks.GroupByTarget()
	for i := range docs {
		fileName := docs[i].FileName()
		linksInDoc, err := docs[i].TikiLinks(docs)
		if err != nil {
			return err
		}
		allOccurrencesLinks := linksToDocs[fileName]
		missingOccurrencesLinks := allOccurrencesLinks.RemoveLinksFromDocs(linksInDoc.ReferencedDocs())
		dedupedOccurrencesLinks := missingOccurrencesLinks.CombineLinksFromSameDocuments()
		dedupedOccurrencesLinks.SortBySourceDocumentTitle()
		newOccurrencesSection, err := RenderOccurrencesSection(dedupedOccurrencesLinks, docs[i])
		if err != nil {
			return fmt.Errorf("error rendering new occurrences sections for document %q: %w", fileName, err)
		}
		existingOccurrencesSection, err := docs[i].FindSectionWithTitle("occurrences")
		if err != nil {
			return fmt.Errorf("error finding existing occurrences sections in document %q: %w", fileName, err)
		}
		var newDoc *domain.Document
		switch {
		case len(dedupedOccurrencesLinks) == 0 && existingOccurrencesSection == nil:
			output.NoChange()
			continue
		case len(dedupedOccurrencesLinks) == 0 && existingOccurrencesSection != nil:
			output.Deleted()
			newDoc = docs[i].RemoveSection(existingOccurrencesSection)
		case len(dedupedOccurrencesLinks) > 0 && existingOccurrencesSection != nil:
			output.Updated()
			newDoc = docs[i].ReplaceSection(existingOccurrencesSection, newOccurrencesSection)
		case len(dedupedOccurrencesLinks) > 0 && existingOccurrencesSection == nil:
			output.Created()
			newDoc = docs[i].AppendSection(newOccurrencesSection)
		}
		err = tikibase.SaveDocument(newDoc)
		if err != nil {
			return fmt.Errorf("cannot update document %s: %w", fileName, err)
		}
	}

	fmt.Println("\n\n" + output.Footer())
	return nil
}
