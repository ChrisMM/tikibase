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
		result, err := check.Run(".")
		if err != nil {
			helpers.PrintErrors(err)
			os.Exit(255)
		}
		handleCheckResults(result)
	},
}

//nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(checkCmd)
}

func handleCheckResults(result check.Result) {
	exitCode := 0

	if len(result.Duplicates) > 0 {
		exitCode += len(result.Duplicates)
		fmt.Printf("\n%d duplicate link targets:\n", len(result.Duplicates))
		for i := range result.Duplicates {
			fmt.Printf("- %s\n", result.Duplicates[i])
		}
		fmt.Println()
	}

	if len(result.BrokenLinks) > 0 {
		exitCode += len(result.BrokenLinks)
		fmt.Printf("\n%d broken links:\n", len(result.BrokenLinks))
		for i := range result.BrokenLinks {
			fmt.Printf("- %s: %q\n", result.BrokenLinks[i].Filename, result.BrokenLinks[i].Link)
		}
	}

	if len(result.NonLinkedResources) > 0 {
		exitCode += len(result.NonLinkedResources)
		fmt.Printf("\n%d non-linked resources:\n", len(result.NonLinkedResources))
		for i := range result.NonLinkedResources {
			fmt.Printf("- %s\n", result.NonLinkedResources[i])
		}
	}

	os.Exit(exitCode)
}
