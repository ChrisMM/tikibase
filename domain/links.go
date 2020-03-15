package domain

// Links is a collection of Links
type Links []Link

// HasLinkTo indicates whether this LinkCollection contains a link that points to the given target.
func (ls Links) HasLinkTo(target string) bool {
	for l := range ls {
		if ls[l].Target() == target {
			return true
		}
	}
	return false
}

// ScaffoldLinks provides LinkCollection instances for testing.
func ScaffoldLinks(targets []string) (result Links) {
	for _, target := range targets {
		result = append(result, ScaffoldLink(target))
	}
	return result
}
