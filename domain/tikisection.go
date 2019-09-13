package domain

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/iancoleman/strcase"

	"github.com/pkg/errors"
)

// TikiSection represents a section in a Document.
type TikiSection struct {
	// the textual content of this TikiSection
	// including the title line
	content TikiSectionContent
}

// TikiSectionContent represents the full content of a TikiSection,
// including heading tag and body.
type TikiSectionContent string

// This global variable is a constant and doesn't need to be stubbed in tests.
//nolint:gochecknoglobals
var markdownLinkRE = regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)

// This global variable is a constant and doesn't need to be stubbed in tests.
//nolint:gochecknoglobals
var htmlLinkRE = regexp.MustCompile(`<a[^>]* href="(.*?)"[^>]*>(.*?)</a>`)

// This global variable is a constant and doesn't need to be stubbed in tests.
//nolint:gochecknoglobals
var stripTitleTagRE = regexp.MustCompile(`#+\s*(.*)`)

// ScaffoldTikiSection provides new TikiSection instances for testing.
func ScaffoldTikiSection(content string) TikiSection {
	return TikiSection{TikiSectionContent(content)}
}

// Anchor provides the URL anchor for this TikiSection.
func (ts TikiSection) Anchor() string {
	return strcase.ToKebab(ts.Title())
}

// Content returns the complete content of the entire section.
func (ts TikiSection) Content() TikiSectionContent {
	return ts.content
}

// TikiLinks returns all TikiLinks in this section.
func (ts TikiSection) TikiLinks(tdc DocumentCollection) (result TikiLinkCollection, err error) {
	matches := markdownLinkRE.FindAllStringSubmatch(string(ts.content), 9999)
	for _, match := range matches {
		linkTitle := match[1]
		targetFileName := DocumentFilename(match[2])
		targetDocument, err := tdc.Find(targetFileName)
		if err != nil {
			return result, errors.Wrapf(err, "cannot find target document ('%s') for link '%s' in Section '%s'", targetFileName, linkTitle, ts.Anchor())
		}
		result = append(result, newTikiLink(linkTitle, ts, targetDocument))
	}
	matches = htmlLinkRE.FindAllStringSubmatch(string(ts.content), 9999)
	fmt.Println(matches)
	for _, match := range matches {
		fmt.Println(match)
		linkTitle := match[2]
		targetFilename := DocumentFilename(match[1])
		targetDocument, err := tdc.Find(targetFilename)
		if err != nil {
			return result, errors.Wrapf(err, "cannot find target document ('%s') for link '%s'", targetFilename, linkTitle)
		}
		result = append(result, newTikiLink(linkTitle, ts, targetDocument))
	}
	return result, nil
}

// Title returns the human-friendly title of this TikiSection,
// i.e. its title tag without the "###"" in front
func (ts TikiSection) Title() string {
	titleLine := strings.SplitN(string(ts.content), "\n", 1)[0]
	matches := stripTitleTagRE.FindStringSubmatch(titleLine)
	return matches[1]
}
