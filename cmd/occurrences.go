package cmd

import (
	"github.com/kevgo/tikibase/helpers"

	"github.com/kevgo/tikibase/occurrences"
	"github.com/spf13/cobra"
)

//nolint:gochecknoglobals
// occurrencesCmd represents the occurrences command
var occurrencesCmd = &cobra.Command{
	Use:   "occurrences",
	Short: "Updates the 'occurrences' section of all documents",
	Long: `Updates the 'occurrences' section of all documents,
which contains backlinks to the current documents.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := occurrences.Run(".")
		helpers.PrintErrors(err)
	},
}

//nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(occurrencesCmd)
}
