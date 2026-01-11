package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gophercloud/gophercloud/v2"
)

// BaseClient represents the base HTTP client for Safir services
type BaseClient struct {
	endpoint      string
	authenticator interface{ GetToken() (string, error) }
	httpClient    *http.Client
	apiVersion    string
	serviceType   ServiceType
}

// BaseClientConfig holds base client configuration
type BaseClientConfig struct {
	Endpoint      string
	Authenticator interface{ GetToken() (string, error) }
	ServiceType   ServiceType
	APIVersion    string
	Timeout       time.Duration
}

// NewBaseClient creates a new base client
func NewBaseClient(config BaseClientConfig) *BaseClient {
	if config.Timeout == 0 {
		config.Timeout = DefaultTimeout
	}

	if config.APIVersion == "" {
		config.APIVersion = DefaultAPIVersion
	}

	// Build full endpoint with API version
	fullEndpoint := BuildEndpointURL(config.Endpoint, config.APIVersion)

	return &BaseClient{
		endpoint:      fullEndpoint,
		authenticator: config.Authenticator,
		httpClient: &http.Client{
			Timeout: config.Timeout,
		},
		apiVersion:  config.APIVersion,
		serviceType: config.ServiceType,
	}
}

// DoRequest performs an HTTP request with automatic token handling
func (c *BaseClient) DoRequest(method, path string, body interface{}) (*http.Response, error) {
	// Get current valid token
	token, err := c.authenticator.GetToken()
	if err != nil {
		return nil, fmt.Errorf("failed to get valid token: %w", err)
	}

	var bodyReader io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewBuffer(jsonData)
	}

	// Build full URL
	url := c.endpoint + path

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("X-Auth-Token", token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "golang-safirclient/1.0")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	// Handle authentication errors
	if resp.StatusCode == http.StatusUnauthorized {
		// Try to re-authenticate if using full authenticator
		if auth, ok := c.authenticator.(*Authenticator); ok && auth.autoReauth {
			resp.Body.Close()
			if err := auth.Reauth(); err != nil {
				return nil, &AuthError{Message: fmt.Sprintf("re-authentication failed: %v", err)}
			}
			// Retry the request with new token
			return c.DoRequest(method, path, body)
		}
		defer resp.Body.Close()
		return nil, &AuthError{Message: "authentication failed: token expired or invalid"}
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		defer resp.Body.Close()
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Message:    string(bodyBytes),
			URL:        req.URL.String(),
			Method:     method,
		}
	}

	return resp, nil
}

// ParseResponse parses JSON response into the provided interface
func (c *BaseClient) ParseResponse(resp *http.Response, v interface{}) error {
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if len(bodyBytes) == 0 {
		// Empty response is valid for some operations (like DELETE)
		return nil
	}

	if err := json.Unmarshal(bodyBytes, v); err != nil {
		return fmt.Errorf("failed to parse response: %w (body: %s)", err, string(bodyBytes))
	}

	return nil
}

// SetTimeout sets the HTTP client timeout
func (c *BaseClient) SetTimeout(timeout time.Duration) {
	c.httpClient.Timeout = timeout
}

// GetTimeout returns the current HTTP client timeout
func (c *BaseClient) GetTimeout() time.Duration {
	return c.httpClient.Timeout
}

// GetEndpoint returns the full API endpoint
func (c *BaseClient) GetEndpoint() string {
	return c.endpoint
}

// GetAPIVersion returns the current API version
func (c *BaseClient) GetAPIVersion() string {
	return c.apiVersion
}

// GetServiceType returns the service type
func (c *BaseClient) GetServiceType() ServiceType {
	return c.serviceType
}

// Ping checks if the service API is accessible
func (c *BaseClient) Ping() error {
	resp, err := c.DoRequest(http.MethodGet, "/", nil)
	if err != nil {
		return fmt.Errorf("ping failed: %w", err)
	}
	resp.Body.Close()
	return nil
}

// GetVersion returns the API version information
func (c *BaseClient) GetVersion() (map[string]interface{}, error) {
	// Remove /v1 from endpoint for version query
	baseEndpoint := strings.TrimSuffix(c.endpoint, "/"+c.apiVersion)

	req, err := http.NewRequest(http.MethodGet, baseEndpoint+"/", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return result, nil
}

// BuildAuthOptions builds AuthOptions from ClientOptions
func BuildAuthOptions(opts ClientOptions) *AuthOptions {
	authOpts := &AuthOptions{
		IdentityEndpoint: opts.AuthURL,
		Username:         opts.Username,
		Password:         opts.Password,
		DomainID:         opts.UserDomainID,
		AllowReauth:      opts.AllowReauth,
	}

	// Build scope
	if opts.ProjectName != "" || opts.ProjectDomainID != "" {
		authOpts.Scope = &gophercloud.AuthScope{
			ProjectName: opts.ProjectName,
			DomainID:    opts.ProjectDomainID,
		}
	}

	return authOpts
}
