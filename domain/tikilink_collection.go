package domain

import (
	"sort"
	"strings"
)

// TikiLinkCollection is a collection of TikiLinks.
type TikiLinkCollection []*TikiLink

// ScaffoldTikiLinkCollection provides TikiLinkCollection instances for testing.
func ScaffoldTikiLinkCollection(data []TikiLinkScaffold) (result TikiLinkCollection) {
	for i := range data {
		result = append(result, ScaffoldTikiLink(data[i]))
	}
	return result
}

// Contains indicates whether this TikiLinkCollection contains the given TikiLink.
func (links TikiLinkCollection) Contains(link *TikiLink) bool {
	for i := range links {
		if *links[i] == *link {
			return true
		}
	}
	return false
}

// CombineLinksFromSameDocuments provides a copy of this TikiLinkCollection
// where multiple links to the same document are combined into one link.
func (links TikiLinkCollection) CombineLinksFromSameDocuments() (result TikiLinkCollection) {
	referencedDocs := make(map[string]TikiLinkCollection)
	for i := range links {
		fileName := links[i].SourceSection().Document().FileName()
		referencedDocs[fileName] = append(referencedDocs[fileName], links[i])
	}
	for _, links := range referencedDocs {
		switch {
		case len(links) == 1:
			result = append(result, links[0])
		case len(links) > 1:
			newLink := newTikiLink("merged link", links[0].SourceSection().Document().TitleSection(), links[0].TargetDocument())
			result = append(result, newLink)
		}
	}
	return result
}

// Filter provides a copy of this TikiLinkCollection
// containing only the elements for which the given filter function is true.
func (links TikiLinkCollection) Filter(filter func(link *TikiLink) bool) (result TikiLinkCollection) {
	for i := range links {
		if filter(links[i]) {
			result = append(result, links[i])
		}
	}
	return result
}

// GroupByTarget determines which links in the given TikiLink list point to which Document.
func (links TikiLinkCollection) GroupByTarget() map[string]TikiLinkCollection {
	result := make(map[string]TikiLinkCollection)
	for i := range links {
		targetFileName := links[i].TargetDocument().FileName()
		result[targetFileName] = append(result[targetFileName], links[i])
	}
	return result
}

// ReferencedDocs provides the Documents that the links in this TikiLinkCollection point to.
func (links TikiLinkCollection) ReferencedDocs() (result DocumentCollection) {
	for i := range links {
		link := links[i]
		doc := link.TargetDocument()
		if !result.Contains(doc) {
			result = append(result, doc)
		}
	}
	return result
}

// RemoveLinksFromDocs provides a copy of this TikiLinkCollection
// that does not contain links from the given Documents.
func (links TikiLinkCollection) RemoveLinksFromDocs(docs DocumentCollection) (result TikiLinkCollection) {
	return links.Filter(func(link *TikiLink) bool {
		return !docs.Contains(link.SourceSection().Document())
	})
}

// SortBySourceDocumentTitle sorts this TikiLinkCollection alphabetically by the target document title.
func (links TikiLinkCollection) SortBySourceDocumentTitle() {
	sort.Slice(links, func(i, j int) bool {
		title1, err := links[i].SourceSection().Document().TitleSection().Title()
		if err != nil {
			panic(err)
		}
		title2, err := links[j].SourceSection().Document().TitleSection().Title()
		if err != nil {
			panic(err)
		}
		return strings.ToLower(title1) < strings.ToLower(title2)
	})
}

// Unique provides a new TikiLinkCollection that contains the links from this one
// with all duplicates removed.
func (links TikiLinkCollection) Unique() (result TikiLinkCollection) {
	for i := range links {
		if !result.Contains(links[i]) {
			result = append(result, links[i])
		}
	}
	return result
}
