package cmd

import "github.com/spf13/cobra"

// instancesCmd represents the instances command
var instancesCmd = &cobra.Command{
	Use:   "instances",
	Short: "Commands for interacting with Cloud Virtual Machine instances.",
	Args:  cobra.NoArgs,
}

func init() {
	rootCmd.AddCommand(instancesCmd)
}
