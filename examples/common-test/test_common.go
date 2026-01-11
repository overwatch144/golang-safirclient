package main

import (
	"log"

	"bitbucket.bilgem.tubitak.gov.tr/scm/hasan.acar/golang-safirclient/common"
	"github.com/gophercloud/gophercloud/v2"
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

	optimizationEndpoint, err := auth.GetEndpoint(common.ServiceTypeOptimization)
	if err != nil {
		log.Printf("⚠ Safir Optimization endpoint not found: %v", err)
	} else {
		log.Printf("✓ Safir Optimization endpoint: %s", optimizationEndpoint)
	}

	migrationEndpoint, err := auth.GetEndpoint(common.ServiceTypeMigration)
	if err != nil {
		log.Printf("⚠ Safir Migration endpoint not found: %v", err)
	} else {
		log.Printf("✓ Safir Migration endpoint: %s", migrationEndpoint)
	}

	cloudWatcherEndpoint, err := auth.GetEndpoint(common.ServiceTypeCloudWatcher)
	if err != nil {
		log.Printf("⚠ Safir Cloud Watcher endpoint not found: %v", err)
	} else {
		log.Printf("✓ Safir Cloud Watcher endpoint: %s", cloudWatcherEndpoint)
	}

	// 3. Test Base Client
	log.Println("\n=== Testing Base Client ===")

	if optimizationEndpoint != "" {
		config := common.BaseClientConfig{
			Endpoint:      optimizationEndpoint,
			Authenticator: auth,
			ServiceType:   common.ServiceTypeOptimization,
			APIVersion:    "v1",
		}

		client := common.NewBaseClient(config)
		log.Printf("✓ Base client created")
		log.Printf("  - Endpoint: %s", client.GetEndpoint())
		log.Printf("  - API Version: %s", client.GetAPIVersion())

		if err := client.Ping(); err != nil {
			log.Printf("⚠ Ping failed: %v", err)
		} else {
			log.Println("✓ Ping successful")
		}
	}

	// 4. Test Utilities
	log.Println("\n=== Testing Utility Functions ===")

	listOpts := &common.ListOptions{
		Limit:   10,
		Marker:  "test-marker",
		SortKey: "created_at",
		SortDir: "desc",
	}
	queryString := common.BuildQueryString(listOpts)
	log.Printf("✓ Query string: %s", queryString)

	log.Println("\n=== All Tests Completed ===")
}
