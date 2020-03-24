package check

// brokenLink describes a broken link.
type brokenLink struct {

	// path of the file containing the broken link
	Filename string

	// the Link that is broken
	Link string
}
