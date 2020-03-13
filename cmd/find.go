package cmd

import (
	"fmt"
	"os"

	"github.com/kevgo/tikibase/find"
	"github.com/spf13/cobra"
)

//nolint:gochecknoglobals
var typeFlag *[]string

//nolint:gochecknoglobals
// findCmd defines the "find" command
var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Searches for entries in this TikiBase",
	Long:  `Searches for entries within the TikiBase, filtered by conditions.`,
	// Args:  cobra.NoArgs(),
	Run: func(cmd *cobra.Command, args []string) {
		results, err := find.Run(".", *typeFlag)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		for _, result := range results {
			fmt.Println(result)
		}
	},
}

//nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(findCmd)
	typeFlag = findCmd.Flags().StringSlice("is", []string{}, "what it is")
}
