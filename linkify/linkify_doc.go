package linkify

import (
	"github.com/kevgo/tikibase/domain"
)

// linkifyDoc replaces all occurrences of the given title
// outside of a link in the given text
// with a linkified version.
func linkifyDoc(doc *domain.Document, docs domain.Documents, docsMappings []docMapping) (result string, err error) {
	// cover all existing links, sections, and URLs in the document text
	docContent := doc.Content()
	existingLinkedDocs, err := doc.LinkedDocs(docs)
	if err != nil {
		return result, err
	}
	existingLinkedDocTitles := []string{}
	for eld := range existingLinkedDocs {
		title, err := existingLinkedDocs[eld].Title()
		if err != nil {
			return result, err
		}
		existingLinkedDocTitles = append(existingLinkedDocTitles, title)
	}

	replacer := NewIgnoringReplacer()
	replacer.Ignore(findLinks(docContent)...)
	replacer.Ignore(findSections(docContent)...)
	replacer.Ignore(findUrls(docContent)...)
	replacer.Ignore(existingLinkedDocTitles...)

	// replace all doc names with a link to the respective doc
	for dm := range docsMappings {
		// don't linkify a doc to itself
		if docsMappings[dm].file == doc.FileName() {
			continue
		}
		// don't add the current mapping if the current doc already contains that link
		if replacer.Ignores(docsMappings[dm].replaceWith) {
			continue
		}
		replacer.ReplaceOnce(docsMappings[dm].lookFor, docsMappings[dm].replaceWith)
	}

	return replacer.Apply(docContent), nil
}
