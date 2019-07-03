package cmd

import "github.com/spf13/cobra"

// zonesCmd represents the regions command
var zonesCmd = &cobra.Command{
	Use:   "zones",
	Short: "Commands for interacting with Tencent Cloud availability zones.",
	Args:  cobra.NoArgs,
}

func init() {
	rootCmd.AddCommand(zonesCmd)
}
