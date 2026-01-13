package optimization

import (
	"fmt"
	"net/http"
)

// ListHostMaintenancePolicies retrieves all host maintenance policies
func (c *Client) ListHostMaintenancePolicies(clusterID *string) ([]HostMaintenancePolicy, error) {
	path := "/host-maintenance"
	if clusterID != nil {
		path = fmt.Sprintf("%s?cluster_id=%s", path, *clusterID)
	}

	resp, err := c.DoRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var policies []HostMaintenancePolicy
	if err := c.ParseResponse(resp, &policies); err != nil {
		return nil, err
	}

	return policies, nil
}

// GetHostMaintenancePolicy retrieves a specific host maintenance policy by ID
func (c *Client) GetHostMaintenancePolicy(policyID string) (*HostMaintenancePolicy, error) {
	path := fmt.Sprintf("/host-maintenance/%s", policyID)
	resp, err := c.DoRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var policy HostMaintenancePolicy
	if err := c.ParseResponse(resp, &policy); err != nil {
		return nil, err
	}

	return &policy, nil
}

// CreateHostMaintenancePolicy creates a new host maintenance policy
func (c *Client) CreateHostMaintenancePolicy(req *HostMaintenancePolicyCreate) (*HostMaintenancePolicy, error) {
	resp, err := c.DoRequest(http.MethodPost, "/host-maintenance", req)
	if err != nil {
		return nil, err
	}

	var response HostMaintenancePolicyCreateResponse
	if err := c.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response.Policy, nil
}

// UpdateHostMaintenancePolicy updates a host maintenance policy
func (c *Client) UpdateHostMaintenancePolicy(policyID string, req *HostMaintenancePolicyUpdate) (*HostMaintenancePolicy, error) {
	path := fmt.Sprintf("/host-maintenance/%s", policyID)
	resp, err := c.DoRequest(http.MethodPut, path, req)
	if err != nil {
		return nil, err
	}

	var response HostMaintenancePolicyUpdateResponse
	if err := c.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response.Policy, nil
}

// DeleteHostMaintenancePolicy deletes a host maintenance policy
func (c *Client) DeleteHostMaintenancePolicy(policyID string) error {
	path := fmt.Sprintf("/host-maintenance/%s", policyID)
	resp, err := c.DoRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}
