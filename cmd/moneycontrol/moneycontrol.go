package moneycontrol

import (
	"github.com/spf13/cobra"
)

// ebayCmd represents the ebay command
var MoneycontrolCmd = &cobra.Command{
	Use:   "moneycontrol",
	Short: "Get Top Gainers",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		moneycontrol(args)
	},
}
