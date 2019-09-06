package storage

import (
	"regexp"

	"github.com/pkg/errors"
)

// TikiSection represents a section in a TikiDocument.
type TikiSection struct {
	// the textual content of this TikiSection
	// including the title line
	content string
}

// NewTikiSection creates a new TikiSection with the given content.
func NewTikiSection(content string) TikiSection {
	return TikiSection{content: content}
}

// Content returns the content of the entire section as a block.
func (ts TikiSection) Content() string {
	return ts.content
}

// TikiLinks returns all TikiLinks in this section.
func (ts TikiSection) TikiLinks(documents *TikiDocumentCollection) ([]TikiLink, error) {
	result := []TikiLink{}
	markdownLinkRE := regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)
	matches := markdownLinkRE.FindAllStringSubmatch(ts.content, 9999)
	for _, match := range matches {
		linkTitle := match[1]
		targetHandle := NewHandleFromFileName(match[2])
		targetDocument, err := documents.Find(targetHandle)
		if err != nil {
			return result, errors.Wrapf(err, "cannot find target document ('%s') for link '%s'", targetHandle, linkTitle)
		}
		result = append(result, NewTikiLink(linkTitle, ts, targetDocument))
	}
	return result, nil
}
