package linkify

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplacer(t *testing.T) {
	text := "Amazon Web Services is a product of Amazon."
	ur := ignoringReplacer{}
	ur.Ignore("Amazon Web Services")
	ur.Replace(regexp.MustCompile("Amazon"), "Jeff Besos")
	replaced := ur.Apply(text)
	assert.Equal(t, "Amazon Web Services is a product of Jeff Besos.", replaced)
}
