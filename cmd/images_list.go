package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/danmrichards/tencent-cloud-cli/internal/tencent"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

var (
	// imagesOffset is the pagination offset.
	imagesOffset uint64

	// imagesLimit is the number of images to return per "page".
	imagesLimit uint64

	// imagesListCmd represents the "images list" command
	imagesListCmd = &cobra.Command{
		Use:   "list",
		Short: "List available images that can be used with Cloud Virtual Machines.",
		Args:  cobra.NoArgs,
		Run:   listImages,
	}
)

func init() {
	imagesCommand.AddCommand(imagesListCmd)

	imagesListCmd.Flags().Uint64Var(
		&imagesOffset,
		"offset",
		0,
		"The pagination offset. Example: If there are 100 items, an offset of 10 and a limit of 20 will return items 11 - 31.",
	)

	imagesListCmd.Flags().Uint64Var(
		&imagesLimit,
		"limit",
		0,
		"The pagination limit. Example: If there are 100 items, an offset of 10 and a limit of 20 will return items 11 - 31.",
	)
}

func listImages(cmd *cobra.Command, _ []string) {
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
	// DescribeImagesRequest can be set. The attribute may be a primitive
	// type or it may reference another data structure. It is recommended to use
	// the IDE for development, which can be easily accessed to view the
	// documentation of each interface and data structure.
	req := cvm.NewDescribeImagesRequest()

	// This interface allows you to set the number of images returned. Specify
	// here to return only one. The SDK uses the pointer style to specify
	// parameters, even for basic types you need to use pointers to assign
	// values to the parameters.
	if imagesOffset > 0 {
		req.Offset = &imagesOffset
	}
	if imagesLimit > 0 {
		req.Limit = &imagesLimit
	}

	res, err := cvmClient.DescribeImages(req)
	if err != nil {
		cmd.PrintErr("Could not get images list: ", err)
		return
	}

	cmd.Println()

	w := tabwriter.NewWriter(os.Stdout, 0, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tName\tStatus\tType\tSize\tOS")

	for _, img := range res.Response.ImageSet {
		fmt.Fprintf(
			w,
			"%s\t%s\t%s\t%s\t%dGB\t%s\n",
			*img.ImageId,
			*img.ImageName,
			*img.ImageState,
			*img.ImageType,
			*img.ImageSize,
			*img.OsName,
		)
	}

	w.Flush()
}
