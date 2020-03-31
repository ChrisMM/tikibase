package linkify

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/kevgo/tikibase/helpers"
)

// replacementLength defines the length of replacement strings.
const replacementLength = 10

// ignoringReplacer replaces given regexes in given text while ignoring given parts of the text.
// Ignores and replacements are case insensitive and applied in the order they were registered.
// The zero value is an empty ignoringReplacer.
type ignoringReplacer struct {
	replacements []replacement
}

// replacement specifies one replacement.
type replacement struct {
	lookFor     *regexp.Regexp
	replaceWith string
	restoreWith string
}

// Apply applies the registered replacements to the given text.
func (c *ignoringReplacer) Apply(text string) string {
	// add all the placeholders
	for r := range c.replacements {
		text = c.replacements[r].lookFor.ReplaceAllLiteralString(text, c.replacements[r].replaceWith)
	}
	// restore all the placeholders
	for r := range c.replacements {
		text = strings.ReplaceAll(text, c.replacements[r].replaceWith, c.replacements[r].restoreWith)
	}
	return text
}

// Ignore makes this replacer ignore the given text when replacing stuff.
func (c *ignoringReplacer) Ignore(term string) {
	c.replacements = append(c.replacements, replacement{
		lookFor:     regexp.MustCompile(fmt.Sprintf(`(?i)%s`, regexp.QuoteMeta(term))),
		replaceWith: helpers.RandomString(replacementLength),
		restoreWith: term,
	})
}

// IgnoreMany makes this replacer ignore the given strings when replacing stuff.
func (c *ignoringReplacer) IgnoreMany(terms []string) {
	for t := range terms {
		c.Ignore(terms[t])
	}
}

// Replace makes this replacer replace all occurrences of the given regex with the given text.
func (c *ignoringReplacer) Replace(re *regexp.Regexp, restoreValue string) {
	c.replacements = append(c.replacements, replacement{
		lookFor:     re,
		replaceWith: helpers.RandomString(replacementLength),
		restoreWith: restoreValue,
	})
}
