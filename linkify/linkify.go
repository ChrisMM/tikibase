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

	// replace all existing linkified versions with placeholders.
	replacements := make(map[string]string)
	existingLinks := FindExistingLinks(text, title)
	for e := range existingLinks {
		replacements[existingLinks[e]] = fmt.Sprintf("{{%s}}", helpers.RandomString(10))
		text = strings.ReplaceAll(text, existingLinks[e], replacements[existingLinks[e]])
	}

	// return if there are no occurrences of title now

	// replace all occurrences of title with a linkified version
	re := regexp.MustCompile(fmt.Sprintf("(?i)%s", title))
	text = re.ReplaceAllString(text, fmt.Sprintf("[%s](%s)", title, target))

	// restore all placeholders
	for r := range replacements {
		text = strings.ReplaceAll(text, replacements[r], r)
	}

	return text
}
