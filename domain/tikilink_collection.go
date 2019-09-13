package domain

// TikiLinkCollection is a collection of TikiLinks.
type TikiLinkCollection []TikiLink

// GroupByTarget determines which links in the given TikiLink list point to which TikiDocument.
func (tlc TikiLinkCollection) GroupByTarget() map[TikiDocumentFilename]TikiLinkCollection {
	result := make(map[TikiDocumentFilename]TikiLinkCollection)
	for _, link := range tlc {
		doc := link.TargetDocument()
		targetFileName := doc.FileName()
		result[targetFileName] = append(result[targetFileName], link)
	}
	return result
}
