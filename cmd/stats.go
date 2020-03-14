package cmd

import (
	"fmt"
	"os"

	"github.com/kevgo/tikibase/helpers"
	"github.com/kevgo/tikibase/stats"
	"github.com/spf13/cobra"
)

//nolint:gochecknoglobals
// statsCmd defines the "stats" command
var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Displays statistics about this TikiBase",
	Long: `Displays statistics about this Tikibase, like:

- the number of documents, sections, links, resources
- section types
`,
	Run: func(cmd *cobra.Command, args []string) {
		result, err := stats.Run(".")
		if err != nil {
			helpers.PrintErrors(err)
			os.Exit(255)
		}
		fmt.Printf("documents: %d\n", result.DocsCount)
		fmt.Printf(" sections: %d\n", result.SectionsCount)
		fmt.Printf("    links: %d\n", result.LinksCount)
		fmt.Printf("resources: %d\n", result.ResourcesCount)
		fmt.Printf("\n%d section types:\n", len(result.SectionTypes))
		for i := range result.SectionTypes {
			fmt.Printf("- %s\n", result.SectionTypes[i])
		}
	},
}

//nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(statsCmd)
}
