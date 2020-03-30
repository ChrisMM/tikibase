package cmd

import (
	"fmt"
	"os"

	"github.com/kevgo/tikibase/helpers"
	"github.com/kevgo/tikibase/linkify"
	"github.com/spf13/cobra"
)

//nolint:gochecknoglobals
// linksCmd represents the occurrences command
var linkifyCmd = &cobra.Command{
	Use:     "linkify",
	Short:   "Finds missing links to other documents",
	Long:    `Finds text passages that look like titles of other documents and makes them links to the respective document.`,
	Aliases: []string{"l"},
	Run: func(cmd *cobra.Command, args []string) {
		logger := func() { fmt.Print(".") }
		err := linkify.Run(".", logger)
		fmt.Println("")
		if err != nil {
			helpers.PrintErrors(err)
			os.Exit(errExitCode)
		}
	},
}

//nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(linkifyCmd)
}
