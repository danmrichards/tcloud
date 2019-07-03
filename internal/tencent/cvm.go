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
// setting a filter.
//
// Specify offset (Offset) and Limit (Limit) to select a part
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

// DescribeInstances is used to query the details of one or more instances.
//
// You can query the details of an instance based on information such as
// instance `ID`, instance name, or instance billing mode. See filter `Filter`
// for details on filtering information. If the parameter is empty, return an
// instance of the current user (the number specified by `Limit`, default is
// 20).
//
// See: https://intl.cloud.tencent.com/document/product/213/15728
func (c *CVMClient) DescribeInstances(req *cvm.DescribeInstancesRequest) (*cvm.DescribeInstancesResponse, error) {
	res, err := c.client.DescribeInstances(req)
	if err != nil {
		if terr, ok := err.(*tcerr.TencentCloudSDKError); ok {
			return nil, fmt.Errorf("api error: %s", terr)
		}
		return nil, err
	}

	return res, nil
}

// RunInstances is used to create an instance of one or more specified
// configurations.
//
// After the instance is created successfully, it will automatically start up,
// and [Instance Status] (/document/api/213/9452#instance_state) becomes
// "Running".
//
// The purchase of the pre-paid instance will pre-deduct the amount required for
// the purchase of this instance. After the hourly payment, the purchase will
// pre-freeze the amount required for the purchase of the instance within one
// hour. Please ensure that the account balance is sufficient before calling
// this interface.
//
// The number of instances allowed to be purchased on this interface follows the
// CVM instance purchase limit, the created instance and the instance share
// quota created by the official website portal.
//
// This interface is an asynchronous interface. When the creation request is
// successfully delivered, it will return an instance `ID` list. At this time,
// the instance is created and immediately completed. During this time, the
// status of the instance will be in the "preparation". You can check the status
// of the corresponding instance by calling the DescribeInstances interface to
// determine the creation. There is no ultimate success. If the status of the
// instance changes from "Ready" to "Running", the creation is successful.
//
// See: https://cloud.tencent.com/document/product/213/2664
// See: https://cloud.tencent.com/document/api/213/15728
func (c *CVMClient) RunInstances(req *cvm.RunInstancesRequest) (*cvm.RunInstancesResponse, error) {
	res, err := c.client.RunInstances(req)
	if err != nil {
		if terr, ok := err.(*tcerr.TencentCloudSDKError); ok {
			return nil, fmt.Errorf("api error: %s", terr)
		}
		return nil, err
	}

	return res, nil
}
