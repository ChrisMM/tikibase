package helpers

import "strings"

// SplitURL splits the given URL into filename and anchor
func SplitURL(link string) (filename string, anchor string) {
	parts := strings.Split(link, "#")
	if len(parts) == 1 {
		return parts[0], ""
	}
	return parts[0], parts[1]
}
