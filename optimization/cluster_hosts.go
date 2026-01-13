package optimization

import (
	"fmt"
	"net/http"
)

// ListClusterHosts retrieves all hosts for a specific cluster
func (c *Client) ListClusterHosts(clusterID string) ([]ClusterHost, error) {
	path := fmt.Sprintf("/clusters/%s/hosts", clusterID)
	resp, err := c.DoRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var hosts []ClusterHost
	if err := c.ParseResponse(resp, &hosts); err != nil {
		return nil, err
	}

	return hosts, nil
}

// GetClusterHost retrieves a specific host by ID
func (c *Client) GetClusterHost(clusterID, hostID string) (*ClusterHost, error) {
	path := fmt.Sprintf("/clusters/%s/hosts/%s", clusterID, hostID)
	resp, err := c.DoRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var host ClusterHost
	if err := c.ParseResponse(resp, &host); err != nil {
		return nil, err
	}

	return &host, nil
}

// CreateClusterHost adds a new host to a cluster
func (c *Client) CreateClusterHost(clusterID string, req *ClusterHostCreate) (*ClusterHost, error) {
	path := fmt.Sprintf("/clusters/%s/hosts", clusterID)
	resp, err := c.DoRequest(http.MethodPost, path, req)
	if err != nil {
		return nil, err
	}

	var response ClusterHostCreateResponse
	if err := c.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response.Host, nil
}

// UpdateClusterHost updates a cluster host
func (c *Client) UpdateClusterHost(clusterID, hostID string, req *ClusterHostUpdate) (*ClusterHost, error) {
	path := fmt.Sprintf("/clusters/%s/hosts/%s", clusterID, hostID)
	resp, err := c.DoRequest(http.MethodPut, path, req)
	if err != nil {
		return nil, err
	}

	var response ClusterHostUpdateResponse
	if err := c.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response.Host, nil
}

// DeleteClusterHost removes a host from a cluster
func (c *Client) DeleteClusterHost(clusterID, hostID string) error {
	path := fmt.Sprintf("/clusters/%s/hosts/%s", clusterID, hostID)
	resp, err := c.DoRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}
