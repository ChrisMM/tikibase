package domain

// LinkCollection is a collection of Links
type LinkCollection []Link

// HasTarget indicates whether this LinkCollection contains a link that points to the given target.
func (lc LinkCollection) HasTarget(target string) bool {
	for l := range lc {
		if lc[l].Target() == target {
			return true
		}
	}
	return false
}

// ScaffoldLinkCollection provides LinkCollection instances for testing.
func ScaffoldLinkCollection(targets []string) (result LinkCollection) {
	for _, target := range targets {
		result = append(result, ScaffoldLink(target))
	}
	return result
}
