package helpers

import (
	"fmt"
	"strings"
)

// PrintErrors prints the given error including all nested errors.
func PrintErrors(err error) {
	errors := strings.Split(err.Error(), ":")
	fmt.Printf("Error: %s:\n", errors[0])
	for _, error := range errors[1:] {
		fmt.Printf("- %s\n", error)
	}
}
