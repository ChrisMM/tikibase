package cmd

import (
	"os"

	"github.com/kevgo/tikibase/helpers"
	"github.com/kevgo/tikibase/remove"
	"github.com/spf13/cobra"
)

//nolint:gochecknoglobals
// removeCmd defines the "find" command
var removeCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Deletes a file",
	Long:    `Deletes a file including all links to this file from this TikiBase.`,
	Aliases: []string{"rm", "d"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := remove.Run(".", args[0])
		if err != nil {
			helpers.PrintErrors(err)
			os.Exit(errExitCode)
		}
	},
}

//nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(removeCmd)
}
