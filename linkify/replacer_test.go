package linkify_test

import (
	"testing"

	"github.com/kevgo/tikibase/linkify"
	"github.com/stretchr/testify/assert"
)

func TestReplacer_Replace(t *testing.T) {
	text := "one two three four"
	r := linkify.NewReplacer()
	r.Add("one", "eins")
	r.Add("three", "drei")
	replaced := r.Replace(text)
	assert.Equal(t, "eins two drei four", replaced)
	assert.Equal(t, text, r.Restore(replaced))
}
