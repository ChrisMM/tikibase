package linkify

import (
	"fmt"
	"regexp"
)

// FindExistingLinks provides all links with the given title in the given text.
func FindExistingLinks(text, title string) (result []string) {
	re := regexp.MustCompile(fmt.Sprintf(`(?i)\[%s\]\(.*?\)`, title))
	matches := re.FindAllString(text, -1)
	collector := make(map[string]struct{})
	for m := range matches {
		collector[matches[m]] = struct{}{}
	}
	for c := range collector {
		result = append(result, c)
	}
	return result
}
