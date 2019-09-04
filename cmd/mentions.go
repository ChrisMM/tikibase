package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// mentionsCmd represents the mentions command
var mentionsCmd = &cobra.Command{
	Use:   "mentions",
	Short: "Updates the 'mentions' section of all documents",
	Long: `Updates the 'mentions' section of all documents.
This section contains backlinks.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mentions called")
	},
}

func init() {
	rootCmd.AddCommand(mentionsCmd)
}
