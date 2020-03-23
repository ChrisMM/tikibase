package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

//nolint:gochecknoglobals
// versionCmd defines the "version" command
var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Displays the TikiBase Tools version",
	Aliases: []string{"v"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TikiBase Tools version 0.0.0")
	},
}

//nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(versionCmd)
}
