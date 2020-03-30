package linkify

import (
	"fmt"
	"regexp"

	"github.com/kevgo/tikibase/domain"
)

// linkifyDoc replaces all occurrences of the given title
// outside of a link in the given text
// with a linkified version.
func linkifyDoc(doc *domain.Document, docsMappings []docMapping) string {
	// cover all existing links, sections, and URLs in the document text
	docContent := doc.Content()

	replacer := ignoringReplacer{}
	replacer.IgnoreMany(findLinks(docContent))
	replacer.IgnoreMany(findSections(docContent))
	replacer.IgnoreMany(findUrls(docContent))

	// replace all doc names with a link to the respective doc
	for dm := range docsMappings {
		// don't linkify a doc to itself
		if docsMappings[dm].file == doc.FileName() {
			continue
		}
		regex := regexp.MustCompile(fmt.Sprintf(`(?i)\b%s\b`, docsMappings[dm].name))
		link := fmt.Sprintf("[%s](%s)", docsMappings[dm].name, docsMappings[dm].file)
		replacer.Replace(regex, link)
	}

	return replacer.Apply(docContent)
}
