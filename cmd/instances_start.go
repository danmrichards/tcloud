package cmd

import (
	"github.com/danmrichards/tcloud/internal/tencent"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

var (
	// instancesStartCmd represents the "instances start" command
	instancesStartCmd = &cobra.Command{
		Use:   "start",
		Short: "Start a Cloud Virtual Machine instance.",
		Example: `# Start an instance
tcloud instances start ins-abcdefgh`,
		Args: cobra.ExactArgs(1),
		Run:  startInstance,
	}
)

func init() {
	instancesCmd.AddCommand(instancesStartCmd)
}

func startInstance(cmd *cobra.Command, args []string) {
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
	// StartInstancesRequest can be set. The attribute may be a primitive type
	// or it may reference another data structure. It is recommended to use the
	// IDE for development, which can be easily accessed to view the
	// documentation of each interface and data structure.
	req := cvm.NewStartInstancesRequest()
	req.InstanceIds = common.StringPtrs([]string{args[0]})

	if _, err := cvmClient.StartInstances(req); err != nil {
		cmd.PrintErr("Could not start instance: ", err)
		return
	}

	cmd.Println("Instance starting")
	cmd.Println(`Check status with "tcloud instances describe ` + args[0] + `"`)
}
