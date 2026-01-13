package optimization

import (
	"fmt"
	"net/http"
)

// ListWorkloadConsolidationPolicies retrieves all workload consolidation policies
func (c *Client) ListWorkloadConsolidationPolicies(clusterID *string) ([]WorkloadConsolidationPolicy, error) {
	path := "/workload-consolidation"
	if clusterID != nil {
		path = fmt.Sprintf("%s?cluster_id=%s", path, *clusterID)
	}

	resp, err := c.DoRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var policies []WorkloadConsolidationPolicy
	if err := c.ParseResponse(resp, &policies); err != nil {
		return nil, err
	}

	return policies, nil
}

// GetWorkloadConsolidationPolicy retrieves a specific workload consolidation policy by ID
func (c *Client) GetWorkloadConsolidationPolicy(policyID string) (*WorkloadConsolidationPolicy, error) {
	path := fmt.Sprintf("/workload-consolidation/%s", policyID)
	resp, err := c.DoRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var policy WorkloadConsolidationPolicy
	if err := c.ParseResponse(resp, &policy); err != nil {
		return nil, err
	}

	return &policy, nil
}

// CreateWorkloadConsolidationPolicy creates a new workload consolidation policy
func (c *Client) CreateWorkloadConsolidationPolicy(req *WorkloadConsolidationPolicyCreate) (*WorkloadConsolidationPolicy, error) {
	resp, err := c.DoRequest(http.MethodPost, "/workload-consolidation", req)
	if err != nil {
		return nil, err
	}

	var response WorkloadConsolidationPolicyCreateResponse
	if err := c.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response.Policy, nil
}

// UpdateWorkloadConsolidationPolicy updates a workload consolidation policy
func (c *Client) UpdateWorkloadConsolidationPolicy(policyID string, req *WorkloadConsolidationPolicyUpdate) (*WorkloadConsolidationPolicy, error) {
	path := fmt.Sprintf("/workload-consolidation/%s", policyID)
	resp, err := c.DoRequest(http.MethodPut, path, req)
	if err != nil {
		return nil, err
	}

	var response WorkloadConsolidationPolicyUpdateResponse
	if err := c.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response.Policy, nil
}

// DeleteWorkloadConsolidationPolicy deletes a workload consolidation policy
func (c *Client) DeleteWorkloadConsolidationPolicy(policyID string) error {
	path := fmt.Sprintf("/workload-consolidation/%s", policyID)
	resp, err := c.DoRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}
