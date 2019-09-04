package cmd

import (
	"github.com/kevgo/tikibase/src/mentions"
	"github.com/spf13/cobra"
)

// mentionsCmd represents the mentions command
var mentionsCmd = &cobra.Command{
	Use:   "mentions",
	Short: "Updates the 'mentions' section of all documents",
	Long: `Updates the 'mentions' section of all documents,
which contains backlinks to the current documents.`,
	Run: func(cmd *cobra.Command, args []string) {
		mentions.Run()
	},
}

func init() {
	rootCmd.AddCommand(mentionsCmd)
}
