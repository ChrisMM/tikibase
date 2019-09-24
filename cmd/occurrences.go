package cmd

import (
	"fmt"
	"os"
	"strings"

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
	rootCmd.AddCommand(occurrencesCmd)
}
