package hashnode

import (
	"github.com/spf13/cobra"
)

var (
	section string = "recent"
)

// ebayCmd represents the ebay command
var HashnodeCmd = &cobra.Command{
	Use:   "hashnode [search term] [flags]",
	Short: "Get hashnode article Topics",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		hashnode(args, section)
	},
}

func init() {
	HashnodeCmd.Flags().StringVarP(&section, "section", "s", "search", "Top, latest, tags, People, Blogs section to scrape")
}
