package occurrences

import (
	"fmt"
	"time"

	"github.com/kevgo/tikibase/domain"
	"github.com/pkg/errors"
)

// Run executes the "occurrences" command in the given directory.
//nolint:funlen
func Run(dir string) error {
	output := Output{startTime: time.Now()}
	tikibase, err := domain.NewTikiBase(dir)
	if err != nil {
		return err
	}
	docs, err := tikibase.Documents()
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
		linksInDoc, err := docs[i].TikiLinks(docs)
		if err != nil {
			return err
		}
		allOccurrencesLinks := linksToDocs[fileName]
		missingOccurrencesLinks := allOccurrencesLinks.RemoveLinksFromDocs(linksInDoc.ReferencedDocs())
		dedupedOccurrencesLinks := missingOccurrencesLinks.CombineLinksFromSameDocuments()
		newOccurrencesSection, err := RenderOccurrencesSection(dedupedOccurrencesLinks, docs[i])
		if err != nil {
			return errors.Wrapf(err, "error rendering new occurrences sections for document '%s'", fileName)
		}
		existingOccurrencesSection, err := docs[i].FindSectionWithTitle("occurrences")
		if err != nil {
			return errors.Wrapf(err, "error finding existing occurrences sections in document '%s'", fileName)
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
			return errors.Wrapf(err, "cannot update document %s", fileName)
		}
	}

	fmt.Println("\n\n" + output.Footer(output.Elapsed(time.Now())))
	return nil
}
