package storage

import (
	"fmt"
	"regexp"

	"github.com/pkg/errors"
)

// TikiSection represents a section in a TikiDocument.
type TikiSection struct {
	// the textual content of this TikiSection
	// including the title line
	content TikiSectionContent
}

// TikiSectionContent represents the full content of a TikiSection,
// including heading tag and body.
type TikiSectionContent string

// This global variables is a constant and doesn't need to be stubbed in tests.
//nolint:gochecknoglobals
var markdownLinkRE = regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)

// This global variables is a constant and doesn't need to be stubbed in tests.
//nolint:gochecknoglobals
var htmlLinkRE = regexp.MustCompile(`<a[^>]* href="(.*?)"[^>]*>(.*?)</a>`)

// NewTikiSection creates a new TikiSection with the given content.
func NewTikiSection(content TikiSectionContent) TikiSection {
	return TikiSection{content: content}
}

// Content returns the complete content of the entire section.
func (ts TikiSection) Content() TikiSectionContent {
	return ts.content
}

// TikiLinks returns all TikiLinks in this section.
func (ts TikiSection) TikiLinks(documents *TikiDocumentCollection) ([]TikiLink, error) {
	result := []TikiLink{}
	matches := markdownLinkRE.FindAllStringSubmatch(string(ts.content), 9999)
	for _, match := range matches {
		linkTitle := match[1]
		targetHandle := NewHandleFromFileName(match[2])
		targetDocument, err := documents.Find(targetHandle)
		if err != nil {
			return result, errors.Wrapf(err, "cannot find target document ('%s') for link '%s'", targetHandle, linkTitle)
		}
		result = append(result, NewTikiLink(linkTitle, ts, targetDocument))
	}
	matches = htmlLinkRE.FindAllStringSubmatch(string(ts.content), 9999)
	fmt.Println(matches)
	for _, match := range matches {
		fmt.Println(match)
		linkTitle := match[2]
		targetHandle := NewHandleFromFileName(match[1])
		targetDocument, err := documents.Find(targetHandle)
		if err != nil {
			return result, errors.Wrapf(err, "cannot find target document ('%s') for link '%s'", targetHandle, linkTitle)
		}
		result = append(result, NewTikiLink(linkTitle, ts, targetDocument))
	}
	return result, nil
}
