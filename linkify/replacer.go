package linkify

import "strings"

// Replacer substitutes text passages in strings
type Replacer struct {
	replacements map[string]string
}

// NewReplacer provides Replacer instances.
func NewReplacer() *Replacer {
	return &Replacer{make(map[string]string)}
}

// Add registers the given replacement for the given term.
func (r *Replacer) Add(term, replacement string) {
	r.replacements[term] = replacement
}

// Replace replaces all registered replacements in the given text.
func (r *Replacer) Replace(text string) string {
	for rr := range r.replacements {
		text = strings.ReplaceAll(text, rr, r.replacements[rr])
	}
	return text
}

// Restore removes the replacements in the given string.
func (r *Replacer) Restore(text string) string {
	for rr := range r.replacements {
		text = strings.ReplaceAll(text, r.replacements[rr], rr)
	}
	return text
}
