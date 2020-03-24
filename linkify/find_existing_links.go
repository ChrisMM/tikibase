package linkify

import (
	"regexp"
	"sort"
)

// findExistingLinks provides all links with the given title in the given text.
func findExistingLinks(text string) (result []string) {
	re := regexp.MustCompile(`\[.*?\]\(.*?\)`)
	matches := re.FindAllString(text, -1)
	collector := make(map[string]struct{})
	for m := range matches {
		collector[matches[m]] = struct{}{}
	}
	for c := range collector {
		result = append(result, c)
	}
	sort.Strings(result)
	return result
}
