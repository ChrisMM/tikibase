package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/kevgo/tikibase/mentions"
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
			errors := strings.Split(err.Error(), ":")
			fmt.Printf("Error: %s:\n", errors[0])
			for _, error := range errors[1:] {
				fmt.Printf("- %s\n", error)
			}
			os.Exit(1)
		}
	},
}

//nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(mentionsCmd)
}
