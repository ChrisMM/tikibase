package linkify

import (
	"regexp"

	"github.com/kevgo/tikibase/helpers"
)

// findLinks provides all links with the given title in the given text.
func findLinks(text string) []string {
	re := regexp.MustCompile(`\[.*?\]\(.*?\)`)
	result := helpers.DedupeStrings(re.FindAllString(text, -1))
	helpers.LongestFirst(result)
	return result
}
