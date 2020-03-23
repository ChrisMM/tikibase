package cmd

import (
	"os"

	"github.com/kevgo/tikibase/check"
	"github.com/kevgo/tikibase/fix"
	"github.com/kevgo/tikibase/helpers"
	"github.com/spf13/cobra"
)

//nolint:gochecknoglobals
// pitstopCmd defines the "pitstop" command
var pitstopCmd = &cobra.Command{
	Use:   "pitstop",
	Short: "Runs all the fixes and checks",
	Long: `Runs everything that can be run after making manual edits to a TikiBase.

This command is equivalent to running "tikibase check" and "tikibase fix".
Please see the documentation of those commands for details on what they do.
`,
	Aliases: []string{"ps"},
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
	rootCmd.AddCommand(pitstopCmd)
}
