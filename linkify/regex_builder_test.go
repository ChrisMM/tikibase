package linkify

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegexBuilder_caseInsensitive(t *testing.T) {
	r := NewRegexBuilder()
	have := r.CaseInsensitive("foo")
	assert.Equal(t, "(?i)foo", have.String())
}

func TestRegexBuilder_caseInsensitiveWholeWord(t *testing.T) {
	r := NewRegexBuilder()
	have := r.CaseInsensitiveWholeWord("foo")
	assert.Equal(t, `(?i)\bfoo\b`, have.String())
}
