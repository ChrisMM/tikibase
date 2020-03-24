package linkify

import (
	"strings"

	"github.com/kevgo/tikibase/helpers"
)

// UniqueReplacer substitutes text passages in text with unique placeholders.
type UniqueReplacer struct {
	replacements map[string]string
}

// NewUniqueReplacer provides Replacer instances.
func NewUniqueReplacer() *UniqueReplacer {
	return &UniqueReplacer{make(map[string]string)}
}

// Add registers the given replacement for the given term.
func (ur *UniqueReplacer) Add(term string) {
	ur.replacements[term] = helpers.RandomString(10)
}

// AddMany registers the given replacement for the given term.
func (ur *UniqueReplacer) AddMany(terms []string) {
	for t := range terms {
		ur.Add(terms[t])
	}
}

// Replace replaces all registered replacements in the given text.
func (ur *UniqueReplacer) Replace(text string) string {
	for rr := range ur.replacements {
		text = strings.ReplaceAll(text, rr, ur.replacements[rr])
	}
	return text
}

// Restore removes the replacements in the given string.
func (ur *UniqueReplacer) Restore(text string) string {
	for rr := range ur.replacements {
		text = strings.ReplaceAll(text, ur.replacements[rr], rr)
	}
	return text
}
