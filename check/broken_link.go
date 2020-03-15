package check

// BrokenLink describes a broken link.
type BrokenLink struct {

	// path of the file containing the broken link
	Filename string

	// the Link that is broken
	Link string
}
