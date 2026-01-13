package optimization

import (
	"fmt"

	"github.com/overwatch144/golang-safirclient/common"
)

// Client represents the Safir Optimization API client
type Client struct {
	*common.BaseClient
}

// ClientOptions represents client configuration options
type ClientOptions struct {
	AuthURL         string
	Username        string
	Password        string
	ProjectName     string
	ProjectDomainID string
	UserDomainID    string
	Region          string
	AllowReauth     bool
}

// NewClient creates a new Safir Optimization client
func NewClient(opts ClientOptions) (*Client, error) {
	// Build auth options
	authOpts := common.BuildAuthOptions(common.ClientOptions{
		AuthURL:         opts.AuthURL,
		Username:        opts.Username,
		Password:        opts.Password,
		ProjectName:     opts.ProjectName,
		ProjectDomainID: opts.ProjectDomainID,
		UserDomainID:    opts.UserDomainID,
		Region:          opts.Region,
		AllowReauth:     opts.AllowReauth,
	})

	// Create authenticator
	auth, err := common.NewAuthenticator(authOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to create authenticator: %w", err)
	}

	// Get Safir Optimization endpoint
	endpoint, err := auth.GetEndpoint(common.ServiceTypeOptimization)
	if err != nil {
		return nil, fmt.Errorf("failed to get optimization endpoint: %w", err)
	}

	// Create base client with /api/v1 prefix
	baseConfig := common.BaseClientConfig{
		Endpoint:      endpoint + "/api",
		Authenticator: auth,
		ServiceType:   common.ServiceTypeOptimization,
		APIVersion:    "v1",
	}

	baseClient := common.NewBaseClient(baseConfig)

	return &Client{
		BaseClient: baseClient,
	}, nil
}

// NewClientWithAuthenticator creates a client with existing authenticator
func NewClientWithAuthenticator(auth *common.Authenticator) (*Client, error) {
	// Get Safir Optimization endpoint
	endpoint, err := auth.GetEndpoint(common.ServiceTypeOptimization)
	if err != nil {
		return nil, fmt.Errorf("failed to get optimization endpoint: %w", err)
	}

	// Create base client with /api/v1 prefix
	baseConfig := common.BaseClientConfig{
		Endpoint:      endpoint + "/api",
		Authenticator: auth,
		ServiceType:   common.ServiceTypeOptimization,
		APIVersion:    "v1",
	}

	baseClient := common.NewBaseClient(baseConfig)

	return &Client{
		BaseClient: baseClient,
	}, nil
}

// NewClientWithToken creates a new Optimization client with existing token and endpoint
// Use this when you already have a valid token and know the endpoint
func NewClientWithToken(endpoint, token string) *Client {
	// Create token authenticator
	tokenAuth := common.NewTokenAuthenticator(endpoint, token)

	// Create base client with /api/v1 prefix
	baseConfig := common.BaseClientConfig{
		Endpoint:      common.NormalizeEndpoint(endpoint) + "/api",
		Authenticator: tokenAuth,
		ServiceType:   common.ServiceTypeOptimization,
		APIVersion:    "v1",
	}

	baseClient := common.NewBaseClient(baseConfig)

	return &Client{
		BaseClient: baseClient,
	}
}
