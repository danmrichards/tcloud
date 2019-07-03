package cmd

import "github.com/spf13/cobra"

// imagesCmd represents the images command
var imagesCmd = &cobra.Command{
	Use:   "images",
	Short: "Commands for interacting with Cloud Virtual Machine images.",
	Args:  cobra.NoArgs,
}

func init() {
	rootCmd.AddCommand(imagesCmd)
}
