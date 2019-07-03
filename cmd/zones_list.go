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

// zonesListCmd represents the "regions list" command
var zonesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available Tencent Cloud availability zones.",
	Example: `# List zones in a region
tcloud zones list --region ap-shanghai
`,
	Args: cobra.NoArgs,
	Run:  listZones,
}

func init() {
	zonesCmd.AddCommand(zonesListCmd)
}

func listZones(cmd *cobra.Command, _ []string) {
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
	// DescribeZonesRequest can be set. The attribute may be a primitive
	// type or it may reference another data structure. It is recommended to use
	// the IDE for development, which can be easily accessed to view the
	// documentation of each interface and data structure.
	req := cvm.NewDescribeZonesRequest()

	res, err := cvmClient.DescribeZones(req)
	if err != nil {
		cmd.PrintErr("Could not get zones list: ", err)
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tStatus")

	for _, z := range res.Response.ZoneSet {
		fmt.Fprintf(
			w,
			"%s\t%s\n",
			*z.Zone,
			*z.ZoneState,
		)
	}

	w.Flush()
}
