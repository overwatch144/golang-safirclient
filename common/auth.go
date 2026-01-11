package common

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack"
	"github.com/gophercloud/gophercloud/v2/openstack/identity/v3/tokens"
)

// Authenticator handles authentication and token management
type Authenticator struct {
	authOptions *AuthOptions
	provider    *gophercloud.ProviderClient
	token       string
	tokenExpiry time.Time
	endpoints   map[ServiceType]string
	mutex       sync.RWMutex
	autoReauth  bool
}

// NewAuthenticator creates a new authenticator instance
func NewAuthenticator(opts *AuthOptions) (*Authenticator, error) {
	if opts == nil {
		return nil, &AuthError{Message: "auth options cannot be nil"}
	}

	// Validate required fields
	if err := ValidateAuthOptions(opts); err != nil {
		return nil, err
	}

	auth := &Authenticator{
		authOptions: opts,
		autoReauth:  opts.AllowReauth,
		endpoints:   make(map[ServiceType]string),
	}

	// Perform initial authentication
	if err := auth.Authenticate(); err != nil {
		return nil, &AuthError{Message: fmt.Sprintf("initial authentication failed: %v", err)}
	}

	return auth, nil
}

// Authenticate performs authentication against Keystone
func (a *Authenticator) Authenticate() error {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	// Build gophercloud auth options
	authOpts := gophercloud.AuthOptions{
		IdentityEndpoint:            a.authOptions.IdentityEndpoint,
		Username:                    a.authOptions.Username,
		UserID:                      a.authOptions.UserID,
		Password:                    a.authOptions.Password,
		DomainID:                    a.authOptions.DomainID,
		DomainName:                  a.authOptions.DomainName,
		TenantID:                    a.authOptions.TenantID,
		TenantName:                  a.authOptions.TenantName,
		AllowReauth:                 a.authOptions.AllowReauth,
		TokenID:                     a.authOptions.TokenID,
		Scope:                       a.authOptions.Scope,
		ApplicationCredentialID:     a.authOptions.ApplicationCredentialID,
		ApplicationCredentialName:   a.authOptions.ApplicationCredentialName,
		ApplicationCredentialSecret: a.authOptions.ApplicationCredentialSecret,
	}

	// Create authenticated client with context
	ctx := context.Background()
	provider, err := openstack.AuthenticatedClient(ctx, authOpts)
	if err != nil {
		return fmt.Errorf("failed to authenticate: %w", err)
	}

	a.provider = provider
	a.token = provider.TokenID

	// Get token expiry time
	if err := a.updateTokenExpiry(); err != nil {
		// Non-fatal error, continue
	}

	// Get service endpoints
	if err := a.discoverEndpoints(); err != nil {
		return fmt.Errorf("failed to discover service endpoints: %w", err)
	}

	return nil
}

// updateTokenExpiry extracts and updates token expiry time
func (a *Authenticator) updateTokenExpiry() error {
	if a.provider == nil {
		return fmt.Errorf("provider not initialized")
	}

	// Create identity v3 client
	identityClient, err := openstack.NewIdentityV3(a.provider, gophercloud.EndpointOpts{})
	if err != nil {
		return fmt.Errorf("failed to create identity client: %w", err)
	}

	// Get token details
	ctx := context.Background()
	tokenDetails, err := tokens.Get(ctx, identityClient, a.token).Extract()
	if err != nil {
		return fmt.Errorf("failed to get token details: %w", err)
	}

	a.tokenExpiry = tokenDetails.ExpiresAt
	return nil
}

// discoverEndpoints discovers all Safir service endpoints
func (a *Authenticator) discoverEndpoints() error {
	if a.provider == nil {
		return fmt.Errorf("provider not initialized")
	}

	// Discover Safir Optimization endpoint
	if endpoint, err := a.getServiceEndpoint(ServiceTypeOptimization); err == nil {
		a.endpoints[ServiceTypeOptimization] = endpoint
	}

	// Discover Safir Migration endpoint
	if endpoint, err := a.getServiceEndpoint(ServiceTypeMigration); err == nil {
		a.endpoints[ServiceTypeMigration] = endpoint
	}

	// Discover Safir Cloud Watcher endpoint
	if endpoint, err := a.getServiceEndpoint(ServiceTypeCloudWatcher); err == nil {
		a.endpoints[ServiceTypeCloudWatcher] = endpoint
	}

	return nil
}

// getServiceEndpoint gets the endpoint for a specific service type
func (a *Authenticator) getServiceEndpoint(serviceType ServiceType) (string, error) {
	endpointOpts := gophercloud.EndpointOpts{
		Type:         string(serviceType),
		Availability: gophercloud.AvailabilityAdmin, // Use admin endpoint by default
	}

	endpoint, err := a.provider.EndpointLocator(endpointOpts)
	if err != nil {
		return "", fmt.Errorf("failed to locate %s endpoint: %w", serviceType, err)
	}

	// Normalize endpoint (remove trailing slashes)
	return NormalizeEndpoint(endpoint), nil
}

// GetToken returns the current valid token
func (a *Authenticator) GetToken() (string, error) {
	a.mutex.RLock()
	token := a.token
	expiry := a.tokenExpiry
	a.mutex.RUnlock()

	// Check if token is expired or about to expire (5 min buffer)
	if !expiry.IsZero() && time.Until(expiry) < 5*time.Minute {
		if a.autoReauth {
			// Token expired or expiring soon, re-authenticate
			if err := a.Authenticate(); err != nil {
				return "", fmt.Errorf("failed to re-authenticate: %w", err)
			}
			// Get new token
			a.mutex.RLock()
			token = a.token
			a.mutex.RUnlock()
		} else {
			return "", &AuthError{Message: "token expired and auto-reauth is disabled"}
		}
	}

	return token, nil
}

// GetEndpoint returns the endpoint for a specific service
func (a *Authenticator) GetEndpoint(serviceType ServiceType) (string, error) {
	a.mutex.RLock()
	defer a.mutex.RUnlock()

	endpoint, exists := a.endpoints[serviceType]
	if !exists {
		return "", fmt.Errorf("endpoint for service %s not found", serviceType)
	}

	return endpoint, nil
}

// GetProvider returns the gophercloud provider client
func (a *Authenticator) GetProvider() *gophercloud.ProviderClient {
	a.mutex.RLock()
	defer a.mutex.RUnlock()
	return a.provider
}

// IsTokenExpired checks if the token is expired
func (a *Authenticator) IsTokenExpired() bool {
	a.mutex.RLock()
	defer a.mutex.RUnlock()

	if a.tokenExpiry.IsZero() {
		return false
	}

	return time.Now().After(a.tokenExpiry)
}

// GetTokenExpiry returns the token expiry time
func (a *Authenticator) GetTokenExpiry() time.Time {
	a.mutex.RLock()
	defer a.mutex.RUnlock()
	return a.tokenExpiry
}

// Reauth forces re-authentication
func (a *Authenticator) Reauth() error {
	return a.Authenticate()
}

// GetAuthInfo returns current authentication information
func (a *Authenticator) GetAuthInfo() AuthInfo {
	a.mutex.RLock()
	defer a.mutex.RUnlock()

	info := AuthInfo{
		Username:    a.authOptions.Username,
		UserID:      a.authOptions.UserID,
		DomainName:  a.authOptions.DomainName,
		DomainID:    a.authOptions.DomainID,
		TokenExpiry: a.tokenExpiry,
	}

	if a.authOptions.Scope != nil {
		info.ProjectName = a.authOptions.Scope.ProjectName
		info.ProjectID = a.authOptions.Scope.ProjectID
	}

	if !a.tokenExpiry.IsZero() {
		info.IsExpired = time.Now().After(a.tokenExpiry)
		info.TimeUntilExpiry = time.Until(a.tokenExpiry)
	}

	return info
}

// TokenAuthenticator creates an authenticator with existing token
type TokenAuthenticator struct {
	endpoint string
	token    string
}

// NewTokenAuthenticator creates authenticator with existing token and endpoint
func NewTokenAuthenticator(endpoint, token string) *TokenAuthenticator {
	return &TokenAuthenticator{
		endpoint: NormalizeEndpoint(endpoint),
		token:    token,
	}
}

// GetToken returns the token (no validation)
func (t *TokenAuthenticator) GetToken() (string, error) {
	if t.token == "" {
		return "", &AuthError{Message: "token is empty"}
	}
	return t.token, nil
}

// GetEndpoint returns the endpoint
func (t *TokenAuthenticator) GetEndpoint() string {
	return t.endpoint
}
