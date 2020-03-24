package linkify

import (
	"fmt"
	"regexp"
)

// Linkify replaces all occurrences of the given title
// outside of a link in the given text
// with a linkified version.
func Linkify(text, title, target string) string {
	// return if there are no occurrences of title
	if !TextContainsTitle(text, title) {
		return text
	}

	// replace all existing links, sections, and URLs
	replacer := NewUniqueReplacer()
	replacer.AddMany(FindExistingLinks(text))
	replacer.AddMany(FindExistingSections(text))
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
