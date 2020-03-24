package linkify

import (
	"fmt"
	"regexp"
	"strings"

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

	// replace all existing links with placeholders.
	replacedText := text
	replacements := make(map[string]string)
	existingLinks := FindExistingLinks(text)
	for e := range existingLinks {
		replacements[existingLinks[e]] = fmt.Sprintf("{{%s}}", helpers.RandomString(10))
		replacedText = strings.ReplaceAll(replacedText, existingLinks[e], replacements[existingLinks[e]])
	}

	// return if there are no occurrences of title now
	if !TextContainsTitle(replacedText, title) {
		return text
	}

	// replace all occurrences of title with a linkified version
	re := regexp.MustCompile(fmt.Sprintf(`(?i)\b%s\b`, title))
	replacedText = re.ReplaceAllString(replacedText, fmt.Sprintf("[%s](%s)", title, target))

	// restore all placeholders
	for r := range replacements {
		replacedText = strings.ReplaceAll(replacedText, replacements[r], r)
	}

	return replacedText
}
