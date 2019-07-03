package cmd

import "github.com/spf13/cobra"

// regionsCmd represents the regions command
var regionsCmd = &cobra.Command{
	Use:   "regions",
	Short: "Commands for interacting with Tencent Cloud regions.",
	Args:  cobra.NoArgs,
}

func init() {
	rootCmd.AddCommand(regionsCmd)
}
