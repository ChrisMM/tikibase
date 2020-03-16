package helpers

import (
	"sort"
	"strings"
)

// SortCaseInsensitive sorts the given string list case insensitive.
func SortCaseInsensitive(list []string) {
	sort.Slice(list, func(i, j int) bool {
		return strings.ToLower(list[i]) < strings.ToLower(list[j])
	})
}
