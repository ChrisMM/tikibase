package linkify

import (
	"regexp"

	"github.com/kevgo/tikibase/helpers"
)

// findExistingLinks provides all links with the given title in the given text.
func findExistingLinks(text string) (result []string) {
	re := regexp.MustCompile(`\[.*?\]\(.*?\)`)
	return helpers.DedupeStrings(re.FindAllString(text, -1))
}
