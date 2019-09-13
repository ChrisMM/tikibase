package domain

// TikiLinkCollection is a collection of TikiLinks.
type TikiLinkCollection []TikiLink

// ScaffoldTikiLinkCollection provides TikiLinkCollection instances for testing.
func ScaffoldTikiLinkCollection(data []TikiLinkScaffold) (result TikiLinkCollection) {
	for _, scaffold := range data {
		result = append(result, ScaffoldTikiLink(scaffold))
	}
	return result
}

// GroupByTarget determines which links in the given TikiLink list point to which Document.
func (tlc TikiLinkCollection) GroupByTarget() map[DocumentFilename]TikiLinkCollection {
	result := make(map[DocumentFilename]TikiLinkCollection)
	for _, link := range tlc {
		doc := link.TargetDocument()
		targetFileName := doc.FileName()
		result[targetFileName] = append(result[targetFileName], link)
	}
	return result
}
