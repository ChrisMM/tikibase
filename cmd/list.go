package cmd

import (
	"fmt"
	"os"

	"github.com/kevgo/tikibase/list"
	"github.com/spf13/cobra"
)

//nolint:gochecknoglobals
var typeFlag *[]string

//nolint:gochecknoglobals
// listCmd defines the "list" command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists entries in this TikiBase",
	Long:  `Lists entries within the TikiBase, filtered by conditions.`,
	// Args:  cobra.NoArgs(),
	Run: func(cmd *cobra.Command, args []string) {
		results, err := list.Run(".", *typeFlag)
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
	rootCmd.AddCommand(listCmd)
	typeFlag = listCmd.Flags().StringSlice("is", []string{}, "what it is")
}
