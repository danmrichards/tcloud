package cmd

import (
	"github.com/danmrichards/tcloud/internal/tencent"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

var (
	// instancesStopCommand represents the "instances stop" command
	instancesStopCommand = &cobra.Command{
		Use:   "stop",
		Short: "Stop a Cloud Virtual Machine instance.",
		Example: `# Stop an instance
tcloud instances stop ins-abcdefgh`,
		Args: cobra.ExactArgs(1),
		Run:  stopInstance,
	}
)

func init() {
	instancesCommand.AddCommand(instancesStopCommand)
}

func stopInstance(cmd *cobra.Command, args []string) {
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
	// StopInstancesRequest can be set. The attribute may be a primitive type or
	// it may reference another data structure. It is recommended to use the IDE
	// for development, which can be easily accessed to view the documentation
	// of each interface and data structure.
	req := cvm.NewStopInstancesRequest()
	req.InstanceIds = common.StringPtrs([]string{args[0]})

	if _, err := cvmClient.StopInstances(req); err != nil {
		cmd.PrintErr("Could not stop instance: ", err)
		return
	}

	cmd.Println("Instance stopping")
	cmd.Println(`Check status with "tcloud instances describe ` + args[0] + `"`)
}
