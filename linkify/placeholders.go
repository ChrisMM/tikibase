package linkify

import "github.com/kevgo/tikibase/helpers"

// replacementLength defines the length of replacement strings.
const replacementLength = 10

// Placeholders stores the placeholders used in a document.
// phrase => placeholder
type Placeholders map[string]string

// NewPlaceholders provides Placeholders instances.
func NewPlaceholders() Placeholders {
	return make(map[string]string)
}

// Make provides the placeholder for the given phrase.
// It creates new ones if needed.
func (p Placeholders) Make(phrase string) string {
	existingPlaceholder, found := p[phrase]
	if found {
		return existingPlaceholder
	}
	p[phrase] = helpers.RandomString(replacementLength)
	return p[phrase]
}
