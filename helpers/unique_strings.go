package helpers

import "sort"

// UniqueStrings removes duplicate elements in the given list of strings.
func UniqueStrings(input []string) (result []string) {
	acc := make(map[string]struct{})
	for i := range input {
		acc[input[i]] = struct{}{}
	}
	result = make([]string, len(acc))
	i := 0
	for a := range acc {
		result[i] = a
		i++
	}
	sort.Strings(result)
	return result
}
