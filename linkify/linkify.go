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

	// find all existing links
	replacedText := text
	replacements := make(map[string]string)
	existingLinks := FindExistingLinks(text)
	for e := range existingLinks {
		replacements[existingLinks[e]] = fmt.Sprintf("{{%s}}", helpers.RandomString(10))
	}

	// replace all section headers with placeholders
	existingSections := FindExistingSections(text)
	for e := range existingSections {
		replacements[existingSections[e]] = fmt.Sprintf("{{%s}}", helpers.RandomString(10))
	}
	for r := range replacements {
		replacedText = strings.ReplaceAll(replacedText, r, replacements[r])
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
