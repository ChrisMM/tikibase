package remove

import (
	"fmt"
	"regexp"
)

// removeLinksToFile provides a copy of the given content with all links to the given file replaces with the link title.
func removeLinksToFile(filename, content string) string {
	linkRE := regexp.MustCompile(fmt.Sprintf(`\[(.*?)\]\(%s\)`, filename))
	return linkRE.ReplaceAllString(content, "$1")
}
