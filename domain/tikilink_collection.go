package domain

// TikiLinkCollection is a collection of TikiLinks.
type TikiLinkCollection []TikiLink

// ScaffoldTikiLinkCollection provides TikiLinkCollection instances for testing.
func ScaffoldTikiLinkCollection(data []TikiLinkScaffold) (result TikiLinkCollection) {
	for i := range data {
		result = append(result, ScaffoldTikiLink(data[i]))
	}
	return result
}

// Contains indicates whether this TikiLinkCollection contains the given TikiLink.
func (tlc TikiLinkCollection) Contains(link TikiLink) bool {
	for i := range tlc {
		if tlc[i].Equal(link) {
			return true
		}
	}
	return false
}

// Equal indicates whether this TikiLinkCollection has the same content as the given TikiLinkCollection.
// This method is needed by https://godoc.org/github.com/google/go-cmp/cmp.
func (tlc TikiLinkCollection) Equal(other TikiLinkCollection) bool {
	if len(tlc) != len(other) {
		return false
	}
	for i := range tlc {
		if tlc[i] != other[i] {
			return false
		}
	}
	return true
}

// GroupByTarget determines which links in the given TikiLink list point to which Document.
func (tlc TikiLinkCollection) GroupByTarget() map[DocumentFilename]TikiLinkCollection {
	result := make(map[DocumentFilename]TikiLinkCollection)
	for i := range tlc {
		targetFileName := tlc[i].TargetDocument().FileName()
		result[targetFileName] = append(result[targetFileName], tlc[i])
	}
	return result
}

// Unique provides a new TikiLinkCollection that contains the links from this one
// with all duplicates removed.
func (tlc TikiLinkCollection) Unique() (result TikiLinkCollection) {
	for i := range tlc {
		link := tlc[i]
		if !result.Contains(link) {
			result = append(result, link)
		}
	}
	return result
}
