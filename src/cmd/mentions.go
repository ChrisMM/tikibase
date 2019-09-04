package cmd

import (
	"log"

	"github.com/kevgo/tikibase/src/mentions"
	"github.com/spf13/cobra"
)

//nolint:gochecknoglobals
// mentionsCmd represents the mentions command
var mentionsCmd = &cobra.Command{
	Use:   "mentions",
	Short: "Updates the 'mentions' section of all documents",
	Long: `Updates the 'mentions' section of all documents,
which contains backlinks to the current documents.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := mentions.Run(".")
		if err != nil {
			log.Fatal(err)
		}
	},
}

//nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(mentionsCmd)
}
