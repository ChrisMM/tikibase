package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

//nolint:gochecknoglobals
// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tikibase",
	Short: "A timeless wiki-like knowledge base",
	Long: `TikiBase is a timeless wiki-like knowledge base with focus on long-term availability.
It's database is a set of human-editable and readable Markdown files.
They can be edited with any Markdown front-end.
TikiBase provides machine-assistance for editing and querying via CLI tools that can also be packaged as bots.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
