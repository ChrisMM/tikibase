package linkify

import (
	"regexp"

	"github.com/kevgo/tikibase/helpers"
)

// findLinks provides all links with the given title in the given text.
func findLinks(text string) (result []string) {
	re := regexp.MustCompile(`\[.*?\]\(.*?\)`)
	return helpers.DedupeStrings(re.FindAllString(text, -1))
}
