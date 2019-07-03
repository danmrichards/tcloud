package cmd

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/danmrichards/tcloud/internal/tencent"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

var (
	// instancesDescribeCommand represents the "instances describe" command
	instancesDescribeCommand = &cobra.Command{
		Use:   "describe",
		Short: "Displays all data associated with a Cloud Virtual Machine.",
		Example: `# Describe an instance
tcloud instances describe foobar`,
		Args: cobra.ExactArgs(1),
		Run:  describeInstance,
	}
)

func init() {
	instancesCommand.AddCommand(instancesDescribeCommand)
}

func describeInstance(cmd *cobra.Command, args []string) {
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
	// DescribeInstancesRequest can be set. The attribute may be a primitive
	// type or it may reference another data structure. It is recommended to use
	// the IDE for development, which can be easily accessed to view the
	// documentation of each interface and data structure.
	req := cvm.NewDescribeInstancesRequest()
	req.InstanceIds = common.StringPtrs([]string{args[0]})

	res, err := cvmClient.DescribeInstances(req)
	if err != nil {
		cmd.PrintErr("Could not get images list: ", err)
		return
	}

	cmd.Println()

	w := tabwriter.NewWriter(os.Stdout, 0, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tName\tStatus\tAvailability Zone\tModel\tIP")

	if len(res.Response.InstanceSet) == 0 {
		return
	}

	in := res.Response.InstanceSet[0]
	fmt.Fprintf(
		w,
		"%s\t%s\t%s\t%s\t%s\t",
		*in.InstanceId,
		*in.InstanceName,
		*in.InstanceState,
		*in.Placement.Zone,
		*in.InstanceType,
	)

	public := make([]string, 0, len(in.PublicIpAddresses))
	for _, pi := range in.PublicIpAddresses {
		if pi == nil {
			continue
		}
		public = append(public, *pi)
	}
	fmt.Fprintf(w, "public: "+strings.Join(public, ", "))
	fmt.Fprintf(w, "\n")

	private := make([]string, 0, len(in.PrivateIpAddresses))
	for _, pr := range in.PrivateIpAddresses {
		if pr == nil {
			continue
		}
		private = append(public, *pr)
	}
	fmt.Fprintf(w, "\t\t\t\t\tprivate: "+strings.Join(private, ", "))
	fmt.Fprintf(w, "\n")

	w.Flush()
}
