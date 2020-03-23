package domain

import (
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/iancoleman/strcase"
	"github.com/kevgo/tikibase/helpers"
)

// Section represents a section in a Document.
type Section struct {
	// the textual content of this TikiSection
	// including the title line
	content string

	// document links to the Document that contains this Section.
	document *Document
}

// SectionScaffold defines the named arguments for ScaffoldSection.
type SectionScaffold struct {
	Content string
	Doc     *Document
}

// This is a global constant that doesn't need to be stubbed in tests.
//nolint:gochecknoglobals
var mdLinkRE *regexp.Regexp

// Helps initialize markdownLinkRE
//nolint:gochecknoglobals
var mdLinkOnce sync.Once

// This global variable is a constant and doesn't need to be stubbed in tests.
//nolint:gochecknoglobals
var htmlImageRE *regexp.Regexp

// Helps initialize htmlLinkRE
//nolint:gochecknoglobals
var htmlImageOnce sync.Once

// This global variable is a constant and doesn't need to be stubbed in tests.
//nolint:gochecknoglobals
var htmlLinkRE *regexp.Regexp

// Helps initialize htmlLinkRE
//nolint:gochecknoglobals
var htmlLinkOnce sync.Once

// This global variable is a constant and doesn't need to be stubbed in tests.
//nolint:gochecknoglobals
var stripTitleTagRE *regexp.Regexp

// Helps initialize stripTitleTagRE
//nolint:gochecknoglobals
var stripTitleTagOnce sync.Once

// ScaffoldSection provides new TikiSection instances for testing.
func ScaffoldSection(data SectionScaffold) *Section {
	if data.Content == "" {
		data.Content = "### default section\ncontent\n"
	}
	if data.Doc == nil {
		data.Doc = ScaffoldDocument(DocumentScaffold{})
	}
	return &Section{content: data.Content, document: data.Doc}
}

// Anchor provides the URL anchor for this TikiSection.
func (section *Section) Anchor() (string, error) {
	sectionTitle, err := section.Title()
	if err != nil {
		return "", fmt.Errorf("cannot determine the anchor of a section: %w\n\n%s", err, section.content)
	}
	return strcase.ToKebab(sectionTitle), nil
}

// AppendText provides a new Section that is this Section with the given line appended.
func (section *Section) AppendText(line string) Section {
	return Section{content: section.content + line, document: section.document}
}

// Content returns the complete content of the entire section.
func (section *Section) Content() string {
	return section.content
}

// Document provides a link to the Document containing this section.
func (section *Section) Document() *Document {
	return section.document
}

// Links returns all links in this section.
func (section *Section) Links() (result []Link) {
	mdLinkOnce.Do(func() { mdLinkRE = regexp.MustCompile(`\[(.*?)\]\((.*?)\)`) })
	matches := mdLinkRE.FindAllStringSubmatch(section.content, 9999)
	for i := range matches {
		title := matches[i][1]
		target := matches[i][2]
		result = append(result, Link{title: title, sourceSection: section, target: target})
	}
	htmlLinkOnce.Do(func() { htmlLinkRE = regexp.MustCompile(`<a[^>]* href="(.*?)"[^>]*>(.*?)</a>`) })
	matches = htmlLinkRE.FindAllStringSubmatch(section.content, 9999)
	for _, match := range matches {
		title := match[2]
		target := match[1]
		result = append(result, Link{title: title, sourceSection: section, target: target})
	}
	htmlImageOnce.Do(func() { htmlImageRE = regexp.MustCompile(`<img[^>]* src="(.*?)"[^>]*>`) })
	matches = htmlImageRE.FindAllStringSubmatch(section.content, 9999)
	for _, match := range matches {
		target := match[1]
		result = append(result, Link{title: "<image>", sourceSection: section, target: target})
	}
	return result
}

// LinkTarget provides the relative URL of this section within this TikiBase.
func (section *Section) LinkTarget() (result string, err error) {
	anchor, err := section.Anchor()
	if err != nil {
		return "", fmt.Errorf("cannot determine the link targets for a section: %w\n\n%s", err, section.content)
	}
	return fmt.Sprintf("%s#%s", section.Document().FileName(), anchor), nil
}

// TikiLinks returns all TikiLinks in this section.
func (section *Section) TikiLinks(dc Documents) (result TikiLinks, err error) {
	sectionTitle, err := section.Title()
	if err != nil {
		return result, err
	}
	for _, link := range section.Links() {
		if helpers.IsURL(link.Target()) {
			// we can ignore links to external files here
			continue
		}
		if strings.HasPrefix(link.Target(), "#") {
			// we can ignore links within the same file here
			continue
		}
		filename, _ := helpers.SplitURL(link.Target())
		if !strings.HasSuffix(filename, ".md") {
			// we can ignore links to non-Markdown files here
			continue
		}
		targetDocument, err := dc.FindByFilename(filename)
		if err != nil {
			return result, fmt.Errorf("cannot find target document (%q) for link %q in Section %q: %w", filename, link.Title(), sectionTitle, err)
		}
		result = append(result, newTikiLink(link.Title(), section, targetDocument))
	}
	return result.Unique(), nil
}

// Title returns the human-friendly title of this TikiSection,
// i.e. its title tag without the "###"" in front
func (section *Section) Title() (string, error) {
	titleLine := strings.SplitN(section.content, "\n", 1)[0]
	stripTitleTagOnce.Do(func() { stripTitleTagRE = regexp.MustCompile(`#+\s*(.*)`) })
	matches := stripTitleTagRE.FindStringSubmatch(titleLine)
	if len(matches) == 0 {
		return "", fmt.Errorf("malformatted section title: %q", titleLine)
	}
	return strings.TrimSpace(matches[1]), nil
}

// URL provides the full URL to this Section (link to document that contains this section + anchor of the section.
func (section *Section) URL() (string, error) {
	sectionAnchor, err := section.Anchor()
	if err != nil {
		return "", fmt.Errorf("cannot determine the URL for section: %w", err)
	}
	return section.document.URL() + "#" + sectionAnchor, nil
}
