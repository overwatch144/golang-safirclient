package common

import (
	"time"

	"github.com/gophercloud/gophercloud/v2"
)

// ClientOptions represents common client configuration options
type ClientOptions struct {
	AuthURL         string
	Username        string
	Password        string
	ProjectName     string
	ProjectDomainID string
	UserDomainID    string
	Region          string
	Timeout         time.Duration
	AllowReauth     bool
}

// AuthOptions contains authentication configuration
type AuthOptions struct {
	IdentityEndpoint            string
	Username                    string
	Password                    string
	UserID                      string
	ApplicationCredentialID     string
	ApplicationCredentialName   string
	ApplicationCredentialSecret string
	DomainID                    string
	DomainName                  string
	TenantID                    string
	TenantName                  string
	AllowReauth                 bool
	TokenID                     string
	Scope                       *gophercloud.AuthScope
}

// Link represents a HATEOAS link
type Link struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
}

// ListOptions represents common list options for pagination
type ListOptions struct {
	Limit   int    `json:"limit,omitempty"`
	Marker  string `json:"marker,omitempty"`
	SortKey string `json:"sort_key,omitempty"`
	SortDir string `json:"sort_dir,omitempty"`
}

// AuthInfo contains authentication information for debugging
type AuthInfo struct {
	Username        string
	UserID          string
	ProjectName     string
	ProjectID       string
	DomainName      string
	DomainID        string
	TokenExpiry     time.Time
	IsExpired       bool
	TimeUntilExpiry time.Duration
}

// ServiceType represents Safir service types
type ServiceType string

const (
	ServiceTypeOptimization ServiceType = "safiroptimization"
	ServiceTypeMigration    ServiceType = "migration"
	ServiceTypeCloudWatcher ServiceType = "cloud_watcher"
)

// String returns the string representation of ServiceType
func (s ServiceType) String() string {
	return string(s)
}

// DefaultTimeout is the default HTTP client timeout
const DefaultTimeout = 30 * time.Second

// DefaultAPIVersion is the default API version for Safir services
const DefaultAPIVersion = "v1"
