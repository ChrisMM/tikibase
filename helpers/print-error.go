package helpers

import (
	"fmt"
	"strings"
)

func PrintErrors(err error) {
	errors := strings.Split(err.Error(), ":")
	fmt.Printf("Error: %s:\n", errors[0])
	for _, error := range errors[1:] {
		fmt.Printf("- %s\n", error)
	}
}
