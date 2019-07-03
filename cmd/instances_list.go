package cmd

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/danmrichards/tcloud/internal/tencent"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

var (
	// instancesOffset is the pagination offset.
	instancesOffset int64

	// instancesLimit is the number of images to return per "page".
	instancesLimit int64

	// instancesListCommand represents the "instances list" command
	instancesListCommand = &cobra.Command{
		Use:   "list",
		Short: "List Cloud Virtual Machine instances.",
		Args:  cobra.NoArgs,
		Run:   listInstances,
	}
)

func init() {
	instancesCommand.AddCommand(instancesListCommand)

	instancesListCommand.Flags().Int64Var(
		&instancesOffset,
		"offset",
		0,
		"The pagination offset. Example: If there are 100 items, an offset of 10 and a limit of 20 will return items 11 - 31.",
	)

	instancesListCommand.Flags().Int64Var(
		&instancesLimit,
		"limit",
		0,
		"The pagination limit. Example: If there are 100 items, an offset of 10 and a limit of 20 will return items 11 - 31.",
	)
}

func listInstances(cmd *cobra.Command, _ []string) {
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

	// This interface allows you to set the number of images returned. Specify
	// here to return only one. The SDK uses the pointer style to specify
	// parameters, even for basic types you need to use pointers to assign
	// values to the parameters.
	if instancesOffset > 0 {
		req.Offset = &instancesOffset
	}
	if instancesLimit > 0 {
		req.Limit = &instancesLimit
	}

	res, err := cvmClient.DescribeInstances(req)
	if err != nil {
		cmd.PrintErr("Could not get images list: ", err)
		return
	}

	cmd.Println()

	w := tabwriter.NewWriter(os.Stdout, 0, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tName\tStatus\tAvailability Zone\tModel\tIP")

	for _, in := range res.Response.InstanceSet {
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
	}

	w.Flush()
}
