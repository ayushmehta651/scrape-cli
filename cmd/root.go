package cmd

import (
	"log"
	"os"

	"github.com/ayushmehta651/scrape-cli/cmd/ebay"
	"github.com/ayushmehta651/scrape-cli/cmd/hashnode"
	"github.com/ayushmehta651/scrape-cli/cmd/moneycontrol"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "scrapy [command] [flags]",
	Short: "Scrape the web from command line",
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func Execute() {
	// Removes the default completion command
	rootCommand.CompletionOptions.DisableDefaultCmd = true
	err := rootCommand.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCommand.AddCommand(ebay.EbayCmd)
	rootCommand.AddCommand(hashnode.HashnodeCmd)
	rootCommand.AddCommand(moneycontrol.MoneycontrolCmd)
}
