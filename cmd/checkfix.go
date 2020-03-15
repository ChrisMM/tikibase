package cmd

import (
	"os"

	"github.com/kevgo/tikibase/check"
	"github.com/kevgo/tikibase/fix"
	"github.com/kevgo/tikibase/helpers"
	"github.com/spf13/cobra"
)

//nolint:gochecknoglobals
// checkfixCmd defines the "checkfix" command
var checkfixCmd = &cobra.Command{
	Use:   "checkfix",
	Short: "Runs all the checks and fixes",
	Long: `Runs everything that can be run after making manual edits to a TikiBase.

This command is equivalent to running "tikibase check" and "tikibase fix".
Please see the documentation of those commands for details on what they do.
`,
	Run: func(cmd *cobra.Command, args []string) {
		err := fix.Run(".")
		if err != nil {
			helpers.PrintErrors(err)
			os.Exit(255)
		}

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
	rootCmd.AddCommand(checkfixCmd)
}
