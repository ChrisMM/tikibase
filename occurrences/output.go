package occurrences

// Output defines the methods to call when signaling output to the CLI.
type Output interface {
	// NoChange gets called if an existing Document was already up to date.
	NoChange()

	// Created gets called if an occurrences section was added to a Document.
	Created()

	// Deleted gets called if an occurrences section was deleted from a Document.
	Deleted()

	// Updated gets called if an occurrences section in a Document was updated.
	Updated()

	// Footer provides the textual summary of an "occurrences" command.
	Footer() string
}
