package optimization

import (
	"fmt"
	"net/http"
)

// ListWorkloadBalancingPolicies retrieves all workload balancing policies
func (c *Client) ListWorkloadBalancingPolicies(clusterID *int) ([]WorkloadBalancingPolicy, error) {
	path := "/workload-balancing"
	if clusterID != nil {
		path = fmt.Sprintf("%s?cluster_id=%d", path, *clusterID)
	}

	resp, err := c.DoRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var policies []WorkloadBalancingPolicy
	if err := c.ParseResponse(resp, &policies); err != nil {
		return nil, err
	}

	return policies, nil
}

// GetWorkloadBalancingPolicy retrieves a specific workload balancing policy by ID
func (c *Client) GetWorkloadBalancingPolicy(policyID int) (*WorkloadBalancingPolicy, error) {
	path := fmt.Sprintf("/workload-balancing/%d", policyID)
	resp, err := c.DoRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var policy WorkloadBalancingPolicy
	if err := c.ParseResponse(resp, &policy); err != nil {
		return nil, err
	}

	return &policy, nil
}

// CreateWorkloadBalancingPolicy creates a new workload balancing policy
func (c *Client) CreateWorkloadBalancingPolicy(req *WorkloadBalancingPolicyCreate) (*WorkloadBalancingPolicy, error) {
	resp, err := c.DoRequest(http.MethodPost, "/workload-balancing", req)
	if err != nil {
		return nil, err
	}

	var response WorkloadBalancingPolicyCreateResponse
	if err := c.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response.Policy, nil
}

// UpdateWorkloadBalancingPolicy updates a workload balancing policy
func (c *Client) UpdateWorkloadBalancingPolicy(policyID int, req *WorkloadBalancingPolicyUpdate) (*WorkloadBalancingPolicy, error) {
	path := fmt.Sprintf("/workload-balancing/%d", policyID)
	resp, err := c.DoRequest(http.MethodPut, path, req)
	if err != nil {
		return nil, err
	}

	var response WorkloadBalancingPolicyUpdateResponse
	if err := c.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response.Policy, nil
}

// DeleteWorkloadBalancingPolicy deletes a workload balancing policy
func (c *Client) DeleteWorkloadBalancingPolicy(policyID int) error {
	path := fmt.Sprintf("/workload-balancing/%d", policyID)
	resp, err := c.DoRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}
