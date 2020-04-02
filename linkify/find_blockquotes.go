package linkify

import (
	"regexp"

	"github.com/kevgo/tikibase/helpers"
)

// findBlockquotes provides all blockquotes in the given text.
func findBlockquotes(text string) []string {
	re := regexp.MustCompile("(?m)^\\s*```\n.*?\n\\s*```$")
	result := helpers.DedupeStrings(re.FindAllString(text, -1))
	return result
}
