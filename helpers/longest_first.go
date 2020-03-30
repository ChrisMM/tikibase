package helpers

import (
	"sort"
	"strings"
)

// LongestFirst sorts the given list of strings by length, then alphabetically.
func LongestFirst(list []string) {
	sort.Slice(list, func(i, j int) bool {
		len1 := len(list[i])
		len2 := len(list[j])
		if len1 != len2 {
			return len1 > len2
		}
		return strings.ToLower(list[i]) < strings.ToLower(list[j])
	})
}
