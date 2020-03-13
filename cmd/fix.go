package cmd

import (
	"fmt"
	"os"

	"github.com/kevgo/tikibase/fix"
	"github.com/kevgo/tikibase/helpers"

	"github.com/spf13/cobra"
)

//nolint:gochecknoglobals
// fixCmd defines the "fix" command
var fixCmd = &cobra.Command{
	Use:   "fix",
	Short: "Fixes all auto-correctable issues in this TikiBase",
	Long: `Fixes all auto-correctable issues in this TikiBase:

- adds an 'occurrences' sections to documents
  containing unmentioned backlinks`,
	Run: func(cmd *cobra.Command, args []string) {
		docsCount, createdCount, updatedCount, deletedCount, err := fix.Run(".")
		if err != nil {
			helpers.PrintErrors(err)
			os.Exit(255)
		}
		fmt.Printf("%d documents, %d created, %d updated, %d deleted\n", docsCount, createdCount, updatedCount, deletedCount)
	},
}

//nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(fixCmd)
}
