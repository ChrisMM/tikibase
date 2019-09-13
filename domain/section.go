package domain

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/iancoleman/strcase"

	"github.com/pkg/errors"
)

// Section represents a section in a Document.
type Section struct {
	// the textual content of this TikiSection
	// including the title line
	content SectionContent
}

// SectionContent represents the full content of a TikiSection,
// including heading tag and body.
type SectionContent string

// This global variable is a constant and doesn't need to be stubbed in tests.
//nolint:gochecknoglobals
var markdownLinkRE = regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)

// This global variable is a constant and doesn't need to be stubbed in tests.
//nolint:gochecknoglobals
var htmlLinkRE = regexp.MustCompile(`<a[^>]* href="(.*?)"[^>]*>(.*?)</a>`)

// This global variable is a constant and doesn't need to be stubbed in tests.
//nolint:gochecknoglobals
var stripTitleTagRE = regexp.MustCompile(`#+\s*(.*)`)

// ScaffoldSection provides new TikiSection instances for testing.
func ScaffoldSection(content string) Section {
	return Section{SectionContent(content)}
}

// Anchor provides the URL anchor for this TikiSection.
func (ts Section) Anchor() string {
	return strcase.ToKebab(ts.Title())
}

// Content returns the complete content of the entire section.
func (ts Section) Content() SectionContent {
	return ts.content
}

// TikiLinks returns all TikiLinks in this section.
func (ts Section) TikiLinks(tdc DocumentCollection) (result TikiLinkCollection, err error) {
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
func (ts Section) Title() string {
	titleLine := strings.SplitN(string(ts.content), "\n", 1)[0]
	matches := stripTitleTagRE.FindStringSubmatch(titleLine)
	return matches[1]
}
