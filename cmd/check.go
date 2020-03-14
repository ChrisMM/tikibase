package cmd

import (
	"fmt"
	"os"

	"github.com/kevgo/tikibase/check"
	"github.com/kevgo/tikibase/helpers"
	"github.com/spf13/cobra"
)

//nolint:gochecknoglobals
// checkCmd defines the "check" command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Checks the consistency of this Tikibase",
	Long: `Runs various checks to ensure the consistency of this Tikibase:

- that links point to existing documents, sections, or files
- that there are no duplicate sections

The exit status indicates the number of problems.
An exit status of 255 indicates an internal error.
Please report internal errors at https://github.com/kevgo/tikibase/issues/new`,
	Run: func(cmd *cobra.Command, args []string) {
		brokenLinks, duplicates, nonLinkedResources, err := check.Run(".")
		if err != nil {
			helpers.PrintErrors(err)
			os.Exit(255)
		}
		exitCode := 0

		if len(duplicates) > 0 {
			exitCode += len(duplicates)
			fmt.Printf("\n%d duplicate link targets:\n", len(duplicates))
			for i := range duplicates {
				fmt.Printf("- %s\n", duplicates[i])
			}
			fmt.Println()
		}

		if len(brokenLinks) > 0 {
			exitCode += len(brokenLinks)
			fmt.Printf("\n%d broken links:\n", len(brokenLinks))
			for i := range brokenLinks {
				fmt.Printf("- %s: %q\n", brokenLinks[i].Filename, brokenLinks[i].Link)
			}
		}

		if len(nonLinkedResources) > 0 {
			exitCode += len(nonLinkedResources)
			fmt.Printf("\n%d non-linked resources:\n", len(nonLinkedResources))
			for i := range nonLinkedResources {
				fmt.Printf("- %s\n", nonLinkedResources[i])
			}
		}

		os.Exit(exitCode)
	},
}

//nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(checkCmd)
}
