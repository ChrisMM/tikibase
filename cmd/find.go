package cmd

import (
	"fmt"
	"os"

	"github.com/kevgo/tikibase/find"
	"github.com/kevgo/tikibase/helpers"
	"github.com/spf13/cobra"
)

//nolint:gochecknoglobals
var typeFlag *[]string

//nolint:gochecknoglobals
// findCmd defines the "find" command
var findCmd = &cobra.Command{
	Use:     "find",
	Short:   "Searches for entries in this TikiBase",
	Long:    `Searches for entries within the TikiBase, filtered by conditions.`,
	Aliases: []string{"fd"},
	Run: func(cmd *cobra.Command, args []string) {
		results, err := find.Run(".", *typeFlag)
		if err != nil {
			helpers.PrintErrors(err)
			os.Exit(errExitCode)
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
