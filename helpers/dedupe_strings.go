package helpers

import (
	"sort"
)

// StringDeduper removes duplicates from a series of strings.
type StringDeduper map[string]struct{}

// NewStringDeduper provides StringDeduper instances.
func NewStringDeduper() *StringDeduper {
	deduper := make(StringDeduper)
	return &deduper
}

// Add accumulates another element of the input series.
func (sd *StringDeduper) Add(text string) {
	(*sd)[text] = struct{}{}
}

// AddMany accumulates many input elements at onc.
func (sd *StringDeduper) AddMany(texts []string) {
	for t := range texts {
		sd.Add(texts[t])
	}
}

// Result provides the deduped input set, unsorted.
func (sd *StringDeduper) Result() []string {
	result := make([]string, len(*sd))
	i := 0
	for a := range *sd {
		result[i] = a
		i++
	}
	return result
}

// Sorted provides the deduped input set, sorted alphabetically.
func (sd *StringDeduper) Sorted() []string {
	result := sd.Result()
	sort.Strings(result)
	return result
}

// SortedCaseInsensitive provides the deduped input set,
// sorted alphabetically and case insensitive.
func (sd *StringDeduper) SortedCaseInsensitive() []string {
	result := sd.Result()
	SortCaseInsensitive(result)
	return result
}

// DedupeStrings removes duplicate elements in the given list of strings.
func DedupeStrings(input []string) (result []string) {
	deduper := NewStringDeduper()
	deduper.AddMany(input)
	return deduper.Sorted()
}
