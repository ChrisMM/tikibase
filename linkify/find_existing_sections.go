package linkify

import (
	"regexp"
	"sort"
	"sync"
)

// This is a global constant that doesn't need to be stubbed in tests.
//nolint:gochecknoglobals
var sectionRE *regexp.Regexp

//nolint:gochecknoglobals
var sectionOnce sync.Once

// FindExistingSections provides the lines containing sections in the given text.
func FindExistingSections(text string) []string {
	sectionOnce.Do(func() { sectionRE = regexp.MustCompile("(?m)^#+.*?\n") })
	hits := sectionRE.FindAllString(text, -1)
	deduper := make(map[string]struct{})
	for h := range hits {
		deduper[hits[h]] = struct{}{}
	}
	result := make([]string, 0, len(deduper))
	for d := range deduper {
		result = append(result, d)
	}
	sort.Strings(result)
	return result
}
