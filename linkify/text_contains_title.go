package linkify

import (
	"regexp"
	"strings"
)

// TextContainsTitle indicates whether the given text contains the given title.
func TextContainsTitle(text, title string) bool {
	re := regexp.MustCompile("(?i)" + strings.Join(strings.Split(title, " "), `\s+`))
	return re.MatchString(text)
}
