package linkify

import (
	"github.com/kevgo/tikibase/domain"
)

// linkifyDoc replaces all occurrences of the given title
// outside of a link in the given text
// with a linkified version.
func linkifyDoc(doc *domain.Document, docsMappings []docMapping) string {
	// cover all existing links, sections, and URLs in the document text
	docContent := doc.Content()

	replacer := NewIgnoringReplacer()
	replacer.Ignore(findLinks(docContent)...)
	replacer.Ignore(findSections(docContent)...)
	replacer.Ignore(findUrls(docContent)...)

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

	return replacer.Apply(docContent)
}
