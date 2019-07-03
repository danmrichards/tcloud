package cmd

import "github.com/spf13/cobra"

// instancesCommand represents the instances command
var instancesCommand = &cobra.Command{
	Use:   "instances",
	Short: "Commands for interacting with Cloud Virtual Machine instances.",
	Args:  cobra.NoArgs,
}

func init() {
	rootCmd.AddCommand(instancesCommand)
}
