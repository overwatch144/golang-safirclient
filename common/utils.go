package common

import (
	"fmt"
	"net/url"
	"strings"
)

// BuildQueryString builds a query string from ListOptions
func BuildQueryString(opts *ListOptions) string {
	if opts == nil {
		return ""
	}

	params := url.Values{}

	if opts.Limit > 0 {
		params.Add("limit", fmt.Sprintf("%d", opts.Limit))
	}

	if opts.Marker != "" {
		params.Add("marker", opts.Marker)
	}

	if opts.SortKey != "" {
		params.Add("sort_key", opts.SortKey)
	}

	if opts.SortDir != "" {
		params.Add("sort_dir", opts.SortDir)
	}

	if len(params) == 0 {
		return ""
	}

	return "?" + params.Encode()
}

// ValidateAuthOptions validates authentication options
func ValidateAuthOptions(opts *AuthOptions) error {
	if opts == nil {
		return &ValidationError{Field: "auth_options", Message: "cannot be nil"}
	}

	if opts.IdentityEndpoint == "" {
		return &ValidationError{Field: "identity_endpoint", Message: "is required"}
	}

	// Check if we have valid authentication method
	hasPassword := opts.Username != "" && opts.Password != ""
	hasToken := opts.TokenID != ""
	hasAppCred := opts.ApplicationCredentialID != "" && opts.ApplicationCredentialSecret != ""
	hasAppCredName := opts.ApplicationCredentialName != "" && opts.ApplicationCredentialSecret != ""

	if !hasPassword && !hasToken && !hasAppCred && !hasAppCredName {
		return &ValidationError{Field: "authentication", Message: "no valid authentication method provided"}
	}

	// Validate scope
	if opts.Scope != nil {
		if opts.Scope.ProjectID == "" && opts.Scope.ProjectName == "" &&
			opts.Scope.DomainID == "" && opts.Scope.DomainName == "" {
			return &ValidationError{Field: "scope", Message: "must specify either project or domain"}
		}
	}

	return nil
}

// NormalizeEndpoint normalizes an endpoint URL by removing trailing slashes
func NormalizeEndpoint(endpoint string) string {
	return strings.TrimRight(endpoint, "/")
}

// BuildEndpointURL builds a complete endpoint URL with version
func BuildEndpointURL(baseEndpoint, version string) string {
	normalized := NormalizeEndpoint(baseEndpoint)
	if version != "" {
		return normalized + "/" + version
	}
	return normalized
}

// GetEnvOrDefault returns environment variable value or default
func GetEnvOrDefault(key, defaultValue string) string {
	// This is a placeholder - in real implementation you'd use os.Getenv
	return defaultValue
}

// Contains checks if a string slice contains a specific string
func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// MergeMaps merges two maps, with values from the second map taking precedence
func MergeMaps(map1, map2 map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	for k, v := range map1 {
		result[k] = v
	}

	for k, v := range map2 {
		result[k] = v
	}

	return result
}
