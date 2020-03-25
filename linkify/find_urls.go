package linkify

import (
	"regexp"
	"sync"

	"mvdan.cc/xurls/v2"
)

// This is a global constant that doesn't need to be stubbed in tests.
//nolint:gochecknoglobals
var urlRE *regexp.Regexp

//nolint:gochecknoglobals
var urlOnce sync.Once

func findUrls(text string) []string {
	urlOnce.Do(func() { urlRE = xurls.Relaxed() })
	return urlRE.FindAllString(text, -1)
}
