package tencent

import (
	"fmt"

	tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

// CVMClient is a wrapper around the Tencent Cloud Virtual Machine API client.
type CVMClient struct {
	client *cvm.Client
}

// CVM returns an instantiated Cloud Virtual Machine API client.
// The CVM client will connect to the API in the given region.
func (a *APIClient) CVM(region string) (*CVMClient, error) {
	// Instantiate the client object to request the product.
	// The second parameter is the geographic information, you can directly fill
	// in the string ap-guangzhou, or reference the preset constant
	c, err := cvm.NewClient(a.credential, region, a.profile)
	if err != nil {
		return nil, err
	}

	return &CVMClient{
		client: c,
	}, nil
}

// DescribeImages is used to view the list of images.
//
// You can query the details of the specified image by specifying the image
// ID, or you can query the details of the image that meets the filter by
// setting a filter. Specify offset (Offset) and Limit (Limit) to select a part
// of the result. By default, the first 20 mirror images that satisfy the
// condition are returned.
//
// See: https://intl.cloud.tencent.com/document/product/213/15715
func (c *CVMClient) DescribeImages(req *cvm.DescribeImagesRequest) (*cvm.DescribeImagesResponse, error) {
	res, err := c.client.DescribeImages(req)
	if err != nil {
		if terr, ok := err.(*tcerr.TencentCloudSDKError); ok {
			return nil, fmt.Errorf("api error: %s", terr)
		}
		return nil, err
	}

	return res, nil
}
