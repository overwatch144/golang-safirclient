package main

import (
	"log"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/yourusername/golang-safirclient/common"
)

func main() {
	log.Println("Testing Golang Safir Client - Common Package")

	// 1. Test Authentication
	log.Println("\n=== Testing Authentication ===")

	authOpts := &common.AuthOptions{
		IdentityEndpoint: "http://10.13.0.10:5000/v3",
		Username:         "admin",
		Password:         "7ea6ee6bfa57e013528cd3bd670ac212831545a38f8764178916fa77fb",
		DomainID:         "default",
		AllowReauth:      true,
		Scope: &gophercloud.AuthScope{
			ProjectName: "admin",
			DomainID:    "default",
		},
	}

	// Validate auth options
	if err := common.ValidateAuthOptions(authOpts); err != nil {
		log.Fatalf("Auth options validation failed: %v", err)
	}
	log.Println("✓ Auth options validated successfully")

	// Create authenticator
	auth, err := common.NewAuthenticator(authOpts)
	if err != nil {
		log.Fatalf("Failed to create authenticator: %v", err)
	}
	log.Println("✓ Authenticator created successfully")

	// Get token
	token, err := auth.GetToken()
	if err != nil {
		log.Fatalf("Failed to get token: %v", err)
	}
	log.Printf("✓ Token obtained: %s...", token[:50])

	// Get auth info
	authInfo := auth.GetAuthInfo()
	log.Printf("✓ Auth Info:")
	log.Printf("  - Username: %s", authInfo.Username)
	log.Printf("  - Project: %s", authInfo.ProjectName)
	log.Printf("  - Token Expiry: %s", authInfo.TokenExpiry)
	log.Printf("  - Time Until Expiry: %s", authInfo.TimeUntilExpiry)
	log.Printf("  - Is Expired: %t", authInfo.IsExpired)

	// 2. Test Service Endpoint Discovery
	log.Println("\n=== Testing Service Endpoint Discovery ===")

	// Try to get Safir Optimization endpoint
	optimizationEndpoint, err := auth.GetEndpoint(common.ServiceTypeOptimization)
	if err != nil {
		log.Printf("⚠ Safir Optimization endpoint not found: %v", err)
	} else {
		log.Printf("✓ Safir Optimization endpoint: %s", optimizationEndpoint)
	}

	// Try to get Safir Migration endpoint
	migrationEndpoint, err := auth.GetEndpoint(common.ServiceTypeMigration)
	if err != nil {
		log.Printf("⚠ Safir Migration endpoint not found: %v", err)
	} else {
		log.Printf("✓ Safir Migration endpoint: %s", migrationEndpoint)
	}

	// Try to get Safir Cloud Watcher endpoint
	cloudWatcherEndpoint, err := auth.GetEndpoint(common.ServiceTypeCloudWatcher)
	if err != nil {
		log.Printf("⚠ Safir Cloud Watcher endpoint not found: %v", err)
	} else {
		log.Printf("✓ Safir Cloud Watcher endpoint: %s", cloudWatcherEndpoint)
	}

	// 3. Test Base Client (if we have an endpoint)
	log.Println("\n=== Testing Base Client ===")

	if optimizationEndpoint != "" {
		config := common.BaseClientConfig{
			Endpoint:      optimizationEndpoint,
			Authenticator: auth,
			ServiceType:   common.ServiceTypeOptimization,
			APIVersion:    "v1",
		}

		client := common.NewBaseClient(config)
		log.Printf("✓ Base client created for Safir Optimization")
		log.Printf("  - Endpoint: %s", client.GetEndpoint())
		log.Printf("  - API Version: %s", client.GetAPIVersion())
		log.Printf("  - Service Type: %s", client.GetServiceType())

		// Test ping
		log.Println("  - Testing ping...")
		if err := client.Ping(); err != nil {
			log.Printf("⚠ Ping failed: %v", err)
		} else {
			log.Println("✓ Ping successful")
		}
	}

	// 4. Test Utility Functions
	log.Println("\n=== Testing Utility Functions ===")

	// Test BuildQueryString
	listOpts := &common.ListOptions{
		Limit:   10,
		Marker:  "test-marker",
		SortKey: "created_at",
		SortDir: "desc",
	}
	queryString := common.BuildQueryString(listOpts)
	log.Printf("✓ Query string: %s", queryString)

	// Test NormalizeEndpoint
	endpoint1 := common.NormalizeEndpoint("http://10.13.0.10:9323/")
	endpoint2 := common.NormalizeEndpoint("http://10.13.0.10:9323")
	log.Printf("✓ Normalized endpoints: %s, %s", endpoint1, endpoint2)

	// Test BuildEndpointURL
	fullURL := common.BuildEndpointURL("http://10.13.0.10:9323", "v1")
	log.Printf("✓ Full endpoint URL: %s", fullURL)

	// 5. Test Error Types
	log.Println("\n=== Testing Error Types ===")

	apiErr := &common.APIError{
		StatusCode: 404,
		Message:    "Resource not found",
		URL:        "http://example.com/resource",
		Method:     "GET",
	}
	log.Printf("✓ APIError: %v", apiErr)
	log.Printf("  - IsNotFound: %t", common.IsNotFound(apiErr))
	log.Printf("  - IsServerError: %t", common.IsServerError(apiErr))

	authErr := &common.AuthError{Message: "Invalid credentials"}
	log.Printf("✓ AuthError: %v", authErr)

	validationErr := &common.ValidationError{
		Field:   "username",
		Message: "cannot be empty",
	}
	log.Printf("✓ ValidationError: %v", validationErr)

	log.Println("\n=== All Tests Completed ===")
}
