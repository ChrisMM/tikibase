package linkify

import (
	"strings"

	"github.com/kevgo/tikibase/helpers"
)

// replacementLength defines the length of replacement strings.
const replacementLength = 10

// uniqueReplacer substitutes text passages in text with unique placeholders.
// Replacements are applied in the order added.
// The zero value is an empty replacer.
type uniqueReplacer struct {
	replacements []replacement
}

// replacement specifies one replacement.
type replacement struct {
	lookFor     string
	replaceWith string
}

// Add registers the given replacement for the given term.
func (ur *uniqueReplacer) Add(term string) {
	ur.replacements = append(ur.replacements, replacement{
		lookFor:     term,
		replaceWith: helpers.RandomString(replacementLength),
	})
}

// AddMany registers the given replacement for the given term.
func (ur *uniqueReplacer) AddMany(terms []string) {
	for t := range terms {
		ur.Add(terms[t])
	}
}

// Replace replaces all registered replacements in the given text.
func (ur *uniqueReplacer) Replace(text string) string {
	for r := range ur.replacements {
		text = strings.ReplaceAll(text, ur.replacements[r].lookFor, ur.replacements[r].replaceWith)
	}
	return text
}

// Restore removes the replacements in the given string.
func (ur *uniqueReplacer) Restore(text string) string {
	for r := range ur.replacements {
		text = strings.ReplaceAll(text, ur.replacements[r].replaceWith, ur.replacements[r].lookFor)
	}
	return text
}
