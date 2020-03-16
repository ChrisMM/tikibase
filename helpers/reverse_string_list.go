package helpers

import "sort"

// ReverseStringList returns the input list reversed.
func ReverseStringList(input []string) {
	sort.Sort(sort.Reverse(sort.StringSlice(input)))
}
