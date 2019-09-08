package domain

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/iancoleman/strcase"

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

// Anchor provides the URL anchor for this TikiSection.
func (ts TikiSection) Anchor() string {
	return strcase.ToKebab(ts.Title())
}

// Content returns the complete content of the entire section.
func (ts TikiSection) Content() TikiSectionContent {
	return ts.content
}

// TikiLinks returns all TikiLinks in this section.
func (ts TikiSection) TikiLinks(tb TikiBase) ([]TikiLink, error) {
	result := []TikiLink{}
	matches := markdownLinkRE.FindAllStringSubmatch(string(ts.content), 9999)
	for _, match := range matches {
		linkTitle := match[1]
		targetFileName := TikiDocumentFilename(match[2])
		targetDocument, err := tb.Load(targetFileName)
		if err != nil {
			return result, errors.Wrapf(err, "cannot find target document ('%s') for link '%s' in Section '%s'", targetFileName, linkTitle, ts.Anchor())
		}
		result = append(result, NewTikiLink(linkTitle, ts, targetDocument))
	}
	matches = htmlLinkRE.FindAllStringSubmatch(string(ts.content), 9999)
	fmt.Println(matches)
	for _, match := range matches {
		fmt.Println(match)
		linkTitle := match[2]
		targetFilename := TikiDocumentFilename(match[1])
		targetDocument, err := tb.Load(targetFilename)
		if err != nil {
			return result, errors.Wrapf(err, "cannot find target document ('%s') for link '%s'", targetFilename, linkTitle)
		}
		result = append(result, NewTikiLink(linkTitle, ts, targetDocument))
	}
	return result, nil
}

// Title returns the human-friendly title of this TikiSection,
// i.e. its title tag without the "###"" in front
func (ts TikiSection) Title() string {
	return strings.SplitN(string(ts.content), "\n", 1)[0]
}
