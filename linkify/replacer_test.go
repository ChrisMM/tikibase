package linkify_test

import (
	"regexp"
	"testing"

	"github.com/kevgo/tikibase/linkify"
	"github.com/stretchr/testify/assert"
)

func TestReplacer_Replace(t *testing.T) {
	text := "one two three four"
	ur := linkify.NewUniqueReplacer()
	ur.AddMany([]string{"one", "three"})
	replaced := ur.Replace(text)
	assert.NotEqual(t, text, replaced)
	assert.Regexp(t, regexp.MustCompile(`\w+ two \w+ four`), replaced)
	assert.Equal(t, text, ur.Restore(replaced))
}