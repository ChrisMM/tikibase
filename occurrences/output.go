package occurrences

import (
	"fmt"
	"strings"
	"time"
)

// Output implements the output to the CLI
type Output struct {
	createdCount int
	updatedCount int
	deletedCount int
	startTime    time.Time
}

// NewOutput provides a new instance of Output,
// initialized to the current time.
func NewOutput() Output {
	return Output{startTime: time.Now()}
}

// ScaffoldOutput provides an Output instance prepopulated with the given counts.
// This is only for testing.
func ScaffoldOutput(created, updated, deleted int) Output {
	return Output{createdCount: created, updatedCount: updated, deletedCount: deleted, startTime: time.Now()}
}

// NoChange gets called if an existing Document was already up to date.
func (o *Output) NoChange() {
}

// Created gets called if an occurrences section was added to a Document.
func (o *Output) Created() {
	o.createdCount++
	fmt.Print(".")
}

// Deleted gets called if an occurrences section was deleted from a Document.
func (o *Output) Deleted() {
	o.deletedCount++
	fmt.Print(".")
}

// Updated gets called if an occurrences section in a Document was updated.
func (o *Output) Updated() {
	o.updatedCount++
	fmt.Print(".")
}

// Footer provides the textual summary of an "occurrences" command.
// To calculate the duration, you can use Output.Elapsed(time.Now()).
func (o *Output) Footer(duration time.Duration) string {
	parts := []string{}
	if o.createdCount == 0 && o.updatedCount == 0 && o.deletedCount == 0 {
		return fmt.Sprintf("no changes, %s", duration)
	}
	if o.createdCount > 0 {
		parts = append(parts, fmt.Sprintf("%d created", o.createdCount))
	}
	if o.updatedCount > 0 {
		parts = append(parts, fmt.Sprintf("%d updated", o.updatedCount))
	}
	if o.deletedCount > 0 {
		parts = append(parts, fmt.Sprintf("%d deleted", o.deletedCount))
	}
	return fmt.Sprintf("%s in %s", strings.Join(parts, ", "), duration)
}

// Elapsed provides the time that has passed between when this output object was created and the given time.
func (o *Output) Elapsed(t time.Time) time.Duration {
	return t.Sub(o.startTime) / time.Millisecond * time.Millisecond
}
