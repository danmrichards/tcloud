package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/danmrichards/tcloud/internal/tencent"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

// regionsListCmd represents the "regions list" command
var regionsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available Tencent Cloud Regions.",
	Example: `# List all regions
tcloud regions list
`,
	Args: cobra.NoArgs,
	Run:  listRegions,
}

func init() {
	regionsCmd.AddCommand(regionsListCmd)
}

func listRegions(cmd *cobra.Command, _ []string) {
	cmd.Println()

	// Note: Many comment blocks in here have been translated from Chinese docs.

	apiClient := tencent.NewAPIClient(
		viper.GetString("tencent_secret_id"),
		viper.GetString("tencent_secret_key"),
		tencent.WithLanguage("en-US"),
	)

	cvmClient, err := apiClient.CVM(apiRegion)
	if err != nil {
		cmd.PrintErr("Could not create Tencent API client: ", err)
		return
	}

	// Instantiate a request object, according to the called interface and the
	// actual situation, you can further set the request parameters. You can
	// directly query the SDK source to determine which properties of
	// DescribeRegionsRequest can be set. The attribute may be a primitive
	// type or it may reference another data structure. It is recommended to use
	// the IDE for development, which can be easily accessed to view the
	// documentation of each interface and data structure.
	req := cvm.NewDescribeRegionsRequest()

	res, err := cvmClient.DescribeRegions(req)
	if err != nil {
		cmd.PrintErr("Could not get regions list: ", err)
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tStatus")

	for _, r := range res.Response.RegionSet {
		fmt.Fprintf(
			w,
			"%s\t%s\n",
			*r.Region,
			*r.RegionState,
		)
	}

	w.Flush()
}
