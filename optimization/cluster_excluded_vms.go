package optimization

import (
	"fmt"
	"net/http"
)

// ListClusterExcludedVMs retrieves all excluded VMs for a specific cluster
func (c *Client) ListClusterExcludedVMs(clusterID string) ([]ClusterExcludedVM, error) {
	path := fmt.Sprintf("/clusters/%s/excluded-vms", clusterID)
	resp, err := c.DoRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var vms []ClusterExcludedVM
	if err := c.ParseResponse(resp, &vms); err != nil {
		return nil, err
	}

	return vms, nil
}

// GetClusterExcludedVM retrieves a specific excluded VM by ID
func (c *Client) GetClusterExcludedVM(clusterID, vmID string) (*ClusterExcludedVM, error) {
	path := fmt.Sprintf("/clusters/%s/excluded-vms/%s", clusterID, vmID)
	resp, err := c.DoRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var vm ClusterExcludedVM
	if err := c.ParseResponse(resp, &vm); err != nil {
		return nil, err
	}

	return &vm, nil
}

// CreateClusterExcludedVM adds a VM to the excluded list
func (c *Client) CreateClusterExcludedVM(clusterID string, req *ClusterExcludedVMCreate) (*ClusterExcludedVM, error) {
	path := fmt.Sprintf("/clusters/%s/excluded-vms", clusterID)
	resp, err := c.DoRequest(http.MethodPost, path, req)
	if err != nil {
		return nil, err
	}

	var response ClusterExcludedVMCreateResponse
	if err := c.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response.VM, nil
}

// DeleteClusterExcludedVM removes a VM from the excluded list
func (c *Client) DeleteClusterExcludedVM(clusterID, vmID string) error {
	path := fmt.Sprintf("/clusters/%s/excluded-vms/%s", clusterID, vmID)
	resp, err := c.DoRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}
