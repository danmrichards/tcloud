package tencent

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

// APIClient is a wrapper around the Tencent Cloud SDK.
type APIClient struct {
	credential *common.Credential
	profile    *profile.ClientProfile
}

// ClientProfileOption is a functional option for modifying attributes of a
// Tencent Cloud client profile.
type ClientProfileOption func(cfp *profile.ClientProfile)

// NewAPIClient returns an instantiated API client.
func NewAPIClient(secretID, secretKey string, opts ...ClientProfileOption) *APIClient {
	// Note: Many comment blocks in here have been translated from Chinese docs.

	// Instantiate an authentication object. The incoming key needs to be sent
	// to the Tencent cloud account key pair secretId, secretKey. This is the
	// way to read from environment variables, you need to set these two values
	// in the environment variable first. You can also write a dead key pair
	// directly in the code, but be careful not to copy, upload or share the
	// code to others. Avoid leaking key pairs that endanger your property.
	creds := common.NewCredential(secretID, secretKey)

	// Instantiate a client configuration object, you can specify the timeout
	// and other configuration.
	cpf := profile.NewClientProfile()

	// Apply options to the client profile.
	for _, opt := range opts {
		opt(cpf)
	}

	return &APIClient{
		credential: creds,
		profile:    cpf,
	}
}

// WithLanguage is a client profile option that sets the profile langauge.
// Valid choices: zh-CN, en-US.
func WithLanguage(lang string) ClientProfileOption {
	return func(cfp *profile.ClientProfile) {
		cfp.Language = lang
	}
}
