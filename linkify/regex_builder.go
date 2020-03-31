package linkify

import (
	"fmt"
	"regexp"
)

// RegexBuilder provides regular expressions efficiently,
// i.e. reusing identical ones.
type RegexBuilder struct {
	cache map[string]*regexp.Regexp
}

// NewRegexBuilder provides instances of RegexBuilder
func NewRegexBuilder() RegexBuilder {
	return RegexBuilder{
		cache: make(map[string]*regexp.Regexp),
	}
}

// CaseInsensitive provides a regex that matches the given phrase case-insensitively.
func (r *RegexBuilder) CaseInsensitive(phrase string) *regexp.Regexp {
	return r.getRegexp(fmt.Sprintf("(?i)%s", regexp.QuoteMeta(phrase)))
}

// CaseInsensitiveWholeWord provides a regex that matches the given phrase as a whole word case-insensitively.
func (r *RegexBuilder) CaseInsensitiveWholeWord(phrase string) *regexp.Regexp {
	return r.getRegexp(fmt.Sprintf(`(?i)\b%s\b`, regexp.QuoteMeta(phrase)))
}

func (r *RegexBuilder) getRegexp(expr string) *regexp.Regexp {
	existingRE, found := r.cache[expr]
	if found {
		return existingRE
	}
	r.cache[expr] = regexp.MustCompile(expr)
	return r.cache[expr]
}
