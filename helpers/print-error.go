package helpers

import (
	"fmt"
	"os"
	"strings"
)

func PrintErrors(err error) {
	if err != nil {
		errors := strings.Split(err.Error(), ":")
		fmt.Printf("Error: %s:\n", errors[0])
		for _, error := range errors[1:] {
			fmt.Printf("- %s\n", error)
		}
		os.Exit(1)
	}
}
