package helpers

import "strings"

// ContainsStrings indicates whether the given haystack contains all given needles.
func ContainsStrings(haystack string, needles []string) bool {
	for _, needle := range needles {
		if !strings.Contains(strings.ToLower(haystack), strings.ToLower(needle)) {
			return false
		}
	}
	return true
}
