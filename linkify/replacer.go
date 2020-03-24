package linkify

import (
	"strings"

	"github.com/kevgo/tikibase/helpers"
)

// uniqueReplacer substitutes text passages in text with unique placeholders.
type uniqueReplacer struct {
	replacements map[string]string
}

// newUniqueReplacer provides Replacer instances.
func newUniqueReplacer() *uniqueReplacer {
	return &uniqueReplacer{make(map[string]string)}
}

// Add registers the given replacement for the given term.
func (ur *uniqueReplacer) Add(term string) {
	ur.replacements[term] = helpers.RandomString(10)
}

// AddMany registers the given replacement for the given term.
func (ur *uniqueReplacer) AddMany(terms []string) {
	for t := range terms {
		ur.Add(terms[t])
	}
}

// Replace replaces all registered replacements in the given text.
func (ur *uniqueReplacer) Replace(text string) string {
	for rr := range ur.replacements {
		text = strings.ReplaceAll(text, rr, ur.replacements[rr])
	}
	return text
}

// Restore removes the replacements in the given string.
func (ur *uniqueReplacer) Restore(text string) string {
	for rr := range ur.replacements {
		text = strings.ReplaceAll(text, ur.replacements[rr], rr)
	}
	return text
}
