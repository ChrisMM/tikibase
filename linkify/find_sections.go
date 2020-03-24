package linkify

import (
	"regexp"
	"sync"

	"github.com/kevgo/tikibase/helpers"
)

// This is a global constant that doesn't need to be stubbed in tests.
//nolint:gochecknoglobals
var sectionRE *regexp.Regexp

//nolint:gochecknoglobals
var sectionOnce sync.Once

// findExistingSections provides the lines containing sections in the given text.
func findSections(text string) []string {
	sectionOnce.Do(func() { sectionRE = regexp.MustCompile("(?m)^#+.*?\n") })
	return helpers.DedupeStrings(sectionRE.FindAllString(text, -1))
}
