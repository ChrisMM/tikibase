package check

import (
	"strings"

	"github.com/kevgo/tikibase/helpers"
)

// FindGroups provides groups of words that differ only in capitalization.
// The given string list must be sorted case-insensitively.
func FindGroups(list []string) (result [][]string) {
	current := 0 // index of the current element
	forward := 0 // for looking forward to see if the next elements are similar to the current one
	for forward < len(list) {
		if current == forward {
			forward++
			continue
		}
		// here, forward is ahead of current
		currentElement := strings.ToLower(list[current])
		forwardElement := strings.ToLower(list[forward])
		if currentElement == forwardElement {
			// the next element is the same as the current element --> try the next one
			forward++
			continue
		}
		// here, forward is ahead and pointing to a different element
		if forward == current+1 {
			// the very next element is different --> no cluster, move on
			current++
			continue
		}
		cluster := helpers.UniqueStrings(list[current:forward])
		helpers.ReverseStringList(cluster)
		result = append(result, cluster)
		current = forward
	}
	// pick up the last cluster
	if forward > current+1 {
		cluster := helpers.UniqueStrings(list[current:forward])
		helpers.ReverseStringList(cluster)
		result = append(result, cluster)
	}
	return result
}
