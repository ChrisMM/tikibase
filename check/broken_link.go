package check

import "github.com/kevgo/tikibase/domain"

// BrokenLink describes a broken link.
type BrokenLink struct {

	// path of the file containing the broken link
	Filename domain.DocumentFilename

	// the Link that is broken
	Link string
}
