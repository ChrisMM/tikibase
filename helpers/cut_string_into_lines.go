package helpers

import "strings"

// CutStringIntoLines cuts the given string into lines,
// including the terminating newlines.
func CutStringIntoLines(text string) (result []string) {
	var builder strings.Builder
	for _, runeValue := range text {
		builder.WriteRune(runeValue)
		if runeValue == '\n' {
			result = append(result, builder.String())
			builder = strings.Builder{}
		}
	}
	if builder.Len() > 0 || len(result) == 0 {
		result = append(result, builder.String())
	}
	return result
}
