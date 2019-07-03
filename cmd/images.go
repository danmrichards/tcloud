package cmd

import "github.com/spf13/cobra"

// imagesCommand represents the images command
var imagesCommand = &cobra.Command{
	Use:   "images",
	Short: "Commands for interacting with Cloud Virtual Machine images.",
	Args:  cobra.NoArgs,
}

func init() {
	rootCmd.AddCommand(imagesCommand)
}
