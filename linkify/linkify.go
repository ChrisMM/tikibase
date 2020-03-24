package linkify

import (
	"fmt"
	"regexp"

	"github.com/kevgo/tikibase/helpers"
)

// Linkify replaces all occurrences of the given title
// outside of a link in the given text
// with a linkified version.
func Linkify(text, title, target string) string {
	// return if there are no occurrences of title
	if !TextContainsTitle(text, title) {
		return text
	}

	// replace all existing links and sections
	replacer := NewReplacer()
	for _, link := range FindExistingLinks(text) {
		replacer.Add(link, fmt.Sprintf("{{%s}}", helpers.RandomString(10)))
	}
	for _, section := range FindExistingSections(text) {
		replacer.Add(section, fmt.Sprintf("{{%s}}", helpers.RandomString(10)))
	}
	replacedText := replacer.Replace(text)

	// return if there are no occurrences of title now
	if !TextContainsTitle(replacedText, title) {
		return text
	}

	// replace all occurrences of title with a linkified version
	re := regexp.MustCompile(fmt.Sprintf(`(?i)\b%s\b`, title))
	replacedText = re.ReplaceAllString(replacedText, fmt.Sprintf("[%s](%s)", title, target))

	// restore all placeholders
	return replacer.Restore(replacedText)
}
