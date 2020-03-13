package cmd

import (
	"fmt"
	"os"

	"github.com/kevgo/tikibase/helpers"

	"github.com/kevgo/tikibase/occurrences"
	"github.com/spf13/cobra"
)

//nolint:gochecknoglobals
// occurrencesCmd defines the "occurrences" command
var occurrencesCmd = &cobra.Command{
	Use:   "occurrences",
	Short: "Updates the 'occurrences' section of all documents",
	Long: `Updates the 'occurrences' section of all documents,
which contains backlinks to the current documents.`,
	Run: func(cmd *cobra.Command, args []string) {
		docsCount, createdCount, updatedCount, deletedCount, err := occurrences.Run(".")
		if err != nil {
			helpers.PrintErrors(err)
			os.Exit(1)
		}
		fmt.Printf("%d documents, %d created, %d updated, %d deleted\n", docsCount, createdCount, updatedCount, deletedCount)
	},
}

//nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(occurrencesCmd)
}
