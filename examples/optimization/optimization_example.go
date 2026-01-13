package main

import (
	"fmt"
	"log"
	"time"

	"github.com/overwatch144/golang-safirclient/optimization"
)

func main() {
	log.Println("=== Safir Optimization Client Test ===\n")

	// Create client
	client, err := optimization.NewClient(optimization.ClientOptions{
		AuthURL:         "http://10.13.0.10:5000/v3",
		Username:        "admin",
		Password:        "7ea6ee6bfa57e013528cd3bd670ac212831545a38f8764178916fa77fb",
		ProjectName:     "admin",
		ProjectDomainID: "default",
		UserDomainID:    "default",
		AllowReauth:     true,
	})
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	log.Println("✓ Client created successfully")

	// Test 1: List Clusters
	log.Println("\n--- Test 1: List Clusters ---")
	clusters, err := client.ListClusters()
	if err != nil {
		log.Printf("⚠ Failed to list clusters: %v", err)
	} else {
		log.Printf("✓ Found %d clusters", len(clusters))
		for _, cluster := range clusters {
			log.Printf("  - %s (ID: %d)", cluster.Name, cluster.ID)
		}
	}

	// Test 2: Create Cluster with unique name
	log.Println("\n--- Test 2: Create Cluster ---")
	clusterName := fmt.Sprintf("Test Cluster %d", time.Now().Unix())
	newCluster, err := client.CreateCluster(&optimization.ClusterCreate{
		Name:        clusterName,
		Description: "Created by golang-safirclient test",
	})
	if err != nil {
		log.Printf("⚠ Failed to create cluster: %v", err)
	} else {
		log.Printf("✓ Cluster created: %s (ID: %d)", newCluster.Name, newCluster.ID)

		// Test 3: Get Cluster
		log.Println("\n--- Test 3: Get Cluster ---")
		cluster, err := client.GetCluster(newCluster.ID)
		if err != nil {
			log.Printf("⚠ Failed to get cluster: %v", err)
		} else {
			log.Printf("✓ Retrieved cluster: %s", cluster.Name)
			log.Printf("  - Description: %s", cluster.Description)
			log.Printf("  - Created At: %s", cluster.CreatedAt)
		}

		// Test 4: Update Cluster
		log.Println("\n--- Test 4: Update Cluster ---")
		updatedCluster, err := client.UpdateCluster(newCluster.ID, &optimization.ClusterUpdate{
			Description: "Updated by golang-safirclient test",
		})
		if err != nil {
			log.Printf("⚠ Failed to update cluster: %v", err)
		} else {
			log.Printf("✓ Cluster updated: %s", updatedCluster.Description)
		}

		// Test 5: Create Cluster Host
		log.Println("\n--- Test 5: Create Cluster Host ---")
		newHost, err := client.CreateClusterHost(newCluster.ID, &optimization.ClusterHostCreate{
			Hostname: "test-host-01.example.com",
			Enabled:  true,
		})
		if err != nil {
			log.Printf("⚠ Failed to create host: %v", err)
		} else {
			log.Printf("✓ Host created: %s (ID: %d)", newHost.Hostname, newHost.ID)

			// Test 6: List Cluster Hosts
			log.Println("\n--- Test 6: List Cluster Hosts ---")
			hosts, err := client.ListClusterHosts(newCluster.ID)
			if err != nil {
				log.Printf("⚠ Failed to list hosts: %v", err)
			} else {
				log.Printf("✓ Found %d hosts", len(hosts))
				for _, host := range hosts {
					log.Printf("  - %s (Enabled: %t)", host.Hostname, host.Enabled)
				}
			}

			// Test 7: Delete Cluster Host
			log.Println("\n--- Test 7: Delete Cluster Host ---")
			if err := client.DeleteClusterHost(newCluster.ID, newHost.ID); err != nil {
				log.Printf("⚠ Failed to delete host: %v", err)
			} else {
				log.Printf("✓ Host deleted successfully")
			}
		}

		// Test 8: Create Excluded VM
		log.Println("\n--- Test 8: Create Excluded VM ---")
		newVM, err := client.CreateClusterExcludedVM(newCluster.ID, &optimization.ClusterExcludedVMCreate{
			VMName: "test-vm-01",
		})
		if err != nil {
			log.Printf("⚠ Failed to create excluded VM: %v", err)
		} else {
			log.Printf("✓ Excluded VM created: %s (ID: %d)", newVM.VMName, newVM.ID)

			// Test 9: List Excluded VMs
			log.Println("\n--- Test 9: List Excluded VMs ---")
			vms, err := client.ListClusterExcludedVMs(newCluster.ID)
			if err != nil {
				log.Printf("⚠ Failed to list excluded VMs: %v", err)
			} else {
				log.Printf("✓ Found %d excluded VMs", len(vms))
				for _, vm := range vms {
					log.Printf("  - %s", vm.VMName)
				}
			}

			// Test 10: Delete Excluded VM
			log.Println("\n--- Test 10: Delete Excluded VM ---")
			if err := client.DeleteClusterExcludedVM(newCluster.ID, newVM.ID); err != nil {
				log.Printf("⚠ Failed to delete excluded VM: %v", err)
			} else {
				log.Printf("✓ Excluded VM deleted successfully")
			}
		}

		// Test 11: Create Host Maintenance Policy
		log.Println("\n--- Test 11: Create Host Maintenance Policy ---")
		newPolicy, err := client.CreateHostMaintenancePolicy(&optimization.HostMaintenancePolicyCreate{
			ClusterID: newCluster.ID,
			Name:      "Test Maintenance Policy",
			Enabled:   true,
		})
		if err != nil {
			log.Printf("⚠ Failed to create host maintenance policy: %v", err)
		} else {
			log.Printf("✓ Policy created: %s (ID: %d)", newPolicy.Name, newPolicy.ID)

			// Cleanup: Delete Policy
			if err := client.DeleteHostMaintenancePolicy(newPolicy.ID); err != nil {
				log.Printf("⚠ Failed to delete policy: %v", err)
			} else {
				log.Printf("✓ Policy deleted successfully")
			}
		}

		// Test 12: Create Workload Balancing Policy
		log.Println("\n--- Test 12: Create Workload Balancing Policy ---")
		balancingPolicy, err := client.CreateWorkloadBalancingPolicy(&optimization.WorkloadBalancingPolicyCreate{
			ClusterID:       newCluster.ID,
			Name:            "Test Balancing Policy",
			BalancingMode:   "moderate",
			CPUBalancing:    true,
			MemoryBalancing: true,
			Period:          3600,
			Enabled:         true,
		})
		if err != nil {
			log.Printf("⚠ Failed to create workload balancing policy: %v", err)
		} else {
			log.Printf("✓ Balancing policy created: %s (Period: %d seconds)", balancingPolicy.Name, balancingPolicy.Period)

			// Cleanup: Delete Policy
			if err := client.DeleteWorkloadBalancingPolicy(balancingPolicy.ID); err != nil {
				log.Printf("⚠ Failed to delete balancing policy: %v", err)
			} else {
				log.Printf("✓ Balancing policy deleted successfully")
			}
		}

		// Test 13: Create Workload Consolidation Policy
		log.Println("\n--- Test 13: Create Workload Consolidation Policy ---")
		consolidationPolicy, err := client.CreateWorkloadConsolidationPolicy(&optimization.WorkloadConsolidationPolicyCreate{
			ClusterID: newCluster.ID,
			Name:      "Test Consolidation Policy",
			Period:    7200,
			Enabled:   true,
		})
		if err != nil {
			log.Printf("⚠ Failed to create workload consolidation policy: %v", err)
		} else {
			log.Printf("✓ Consolidation policy created: %s (Period: %d seconds)", consolidationPolicy.Name, consolidationPolicy.Period)

			// Cleanup: Delete Policy
			if err := client.DeleteWorkloadConsolidationPolicy(consolidationPolicy.ID); err != nil {
				log.Printf("⚠ Failed to delete consolidation policy: %v", err)
			} else {
				log.Printf("✓ Consolidation policy deleted successfully")
			}
		}

		// Cleanup: Delete Cluster
		log.Println("\n--- Cleanup: Delete Cluster ---")
		if err := client.DeleteCluster(newCluster.ID); err != nil {
			log.Printf("⚠ Failed to delete cluster: %v", err)
		} else {
			log.Printf("✓ Cluster deleted successfully")
		}
	}

	log.Println("\n=== All Tests Completed ===")
}
