package cmd

import (
	"github.com/danmrichards/tcloud/internal/tencent"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

var (
	// instancesCreateCmd represents the "instances create" command
	instancesCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Creates a new Cloud Virtual Machine instance.",
		Example: `# Create an instance
tcloud instances describe create ins-abcdefgh --image img-abcde`,
		Args: cobra.ExactArgs(1),
		Run:  createInstance,
	}

	// imageID is the image used for the created Compute Virtual Machine.
	imageID string

	// zone is the availability zone in which the Compute Virtual Machine will
	// be created.
	zone string
)

func init() {
	instancesCmd.AddCommand(instancesCreateCmd)

	instancesCreateCmd.Flags().StringVar(
		&imageID,
		"image",
		"",
		"ID of the image used for the created Compute Virtual Machine.",
	)
	instancesCreateCmd.MarkFlagRequired("image")

	instancesCreateCmd.Flags().StringVar(
		&zone,
		"zone",
		"",
		"Availability zone in which the Compute Virtual Machine will be created. See: tcloud zones list.",
	)
	instancesCreateCmd.MarkFlagRequired("zone")
}

func createInstance(cmd *cobra.Command, args []string) {
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
	// RunInstancesRequest can be set. The attribute may be a primitive type or
	// it may reference another data structure. It is recommended to use the IDE
	// for development, which can be easily accessed to view the documentation
	// of each interface and data structure.
	req := cvm.NewRunInstancesRequest()

	// Location of the instance. This parameter is used to specify the
	// availability zone, the project to which the instance belongs, and the
	// CDH (the creation of sub-machine for the exclusive parent host billing
	// mode), etc.
	req.Placement = &cvm.Placement{
		Zone: &zone,
	}

	// The display name of the instance. Not actually required by the API but if
	// it is not specified, "Not named" is displayed.
	req.InstanceName = &args[0]

	// Specify a valid Image ID, in the format of img-xxx. There are four types
	// of images:
	//   Public image
	//   Custom image
	//   Shared image
	//   Marketplace image
	//
	// You can obtain the available image IDs by either of the following ways:
	// * Query the image ID of a public image, custom image or shared image by
	//   logging in to the Console; query the image ID of a marketplace image
	//   via Cloud Marketplace.
	// * Call the API DescribeImages to obtain the ImageId field value in the
	//   returned result.
	req.ImageId = &imageID

	res, err := cvmClient.RunInstances(req)
	if err != nil {
		cmd.PrintErr("Could not create instance: ", err)
		return
	}
	instanceID := *res.Response.InstanceIdSet[0]

	cmd.Printf("Instance creating with ID %q\n", instanceID)
	cmd.Println(`Check status with "tcloud instances describe ` + instanceID + `"`)
}
