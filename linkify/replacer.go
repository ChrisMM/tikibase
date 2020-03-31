package linkify

import (
	"strings"

	"github.com/kevgo/tikibase/helpers"
)

// IgnoringReplacer replaces given phrases in text while ignoring parts of the text.
// Ignores and replacements are case insensitive and applied longest first.
// The zero value is an empty ignoringReplacer.
type IgnoringReplacer struct {
	ignore      []string          // phrases to ignore
	replaceOnce map[string]string // phrases to replace once (old => new)
	regexes     RegexBuilder
}

// NewIgnoringReplacer provides instances of ignoringReplacer.
func NewIgnoringReplacer() IgnoringReplacer {
	return IgnoringReplacer{
		replaceOnce: map[string]string{},
		regexes:     NewRegexBuilder(),
	}
}

// Apply applies the registered replacements to the given text.
func (c *IgnoringReplacer) Apply(text string) string {
	// To prevent replacing terms inside longer terms (i.e. replace "Amazon" inside "Amazon Web Services"),
	// this method temporarily "inks out" (replaces with placeholders) all areas that it has processed,
	// and restores the placeholders when done.
	// To achieve robustness for all edge cases, all replacements are done this way.
	// To ensure the replacements are as specific as possible, this method starts replacing the longer phrases first.
	placeholders := NewPlaceholders()

	// ink out all the phrases to ignore (longest first)
	helpers.LongestFirst(c.ignore)
	for i := range c.ignore {
		re := c.regexes.CaseInsensitive(c.ignore[i])
		text = re.ReplaceAllStringFunc(text, placeholders.Make)
	}

	// ink out the replacements (longest first)
	replacementPhrases := c.phrasesToReplace()
	for r := range replacementPhrases {
		re := c.regexes.CaseInsensitiveWholeWord(replacementPhrases[r])
		hasReplaced := false
		text = re.ReplaceAllStringFunc(text, func(match string) string {
			if !hasReplaced {
				// first occurrence: ink out with the actual replacement
				hasReplaced = true
				return placeholders.Make(c.replaceOnce[replacementPhrases[r]])
			}
			// other occurrences: ink out with the match
			return placeholders.Make(match)
		})
	}

	// restore all the placeholders
	for phrase := range placeholders {
		text = strings.ReplaceAll(text, placeholders[phrase], phrase)
	}
	return text
}

// Ignore makes this replacer ignore the given text when replacing stuff.
func (c *IgnoringReplacer) Ignore(phrase ...string) {
	c.ignore = append(c.ignore, phrase...)
}

// Ignores indicates whether this replacer already ignores the given phrase (case insensitively).
func (c *IgnoringReplacer) Ignores(phrase string) bool {
	for i := range c.ignore {
		if strings.EqualFold(c.ignore[i], phrase) {
			return true
		}
	}
	return false
}

// phrasesToReplace provides all phrases that should be replaced, sorted by length.
func (c *IgnoringReplacer) phrasesToReplace() []string {
	result := make([]string, 0, len(c.replaceOnce))
	for phrase := range c.replaceOnce {
		result = append(result, phrase)
	}
	helpers.LongestFirst(result)
	return result
}

// ReplaceOnce replaces the first occurrence of old with new,
// the remaining occurrences are ignored.
func (c *IgnoringReplacer) ReplaceOnce(old string, new string) {
	c.replaceOnce[old] = new
}
