package optimization

import (
	"fmt"
	"net/http"
)

// ListClusters retrieves all clusters
func (c *Client) ListClusters() ([]Cluster, error) {
	resp, err := c.DoRequest(http.MethodGet, "/clusters", nil)
	if err != nil {
		return nil, err
	}

	var clusters []Cluster
	if err := c.ParseResponse(resp, &clusters); err != nil {
		return nil, err
	}

	return clusters, nil
}

// GetCluster retrieves a specific cluster by ID
func (c *Client) GetCluster(clusterID int) (*Cluster, error) {
	path := fmt.Sprintf("/clusters/%d", clusterID)
	resp, err := c.DoRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var cluster Cluster
	if err := c.ParseResponse(resp, &cluster); err != nil {
		return nil, err
	}

	return &cluster, nil
}

// CreateCluster creates a new cluster
func (c *Client) CreateCluster(req *ClusterCreate) (*Cluster, error) {
	resp, err := c.DoRequest(http.MethodPost, "/clusters", req)
	if err != nil {
		return nil, err
	}

	var response ClusterCreateResponse
	if err := c.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response.Cluster, nil
}

// UpdateCluster updates an existing cluster
func (c *Client) UpdateCluster(clusterID int, req *ClusterUpdate) (*Cluster, error) {
	path := fmt.Sprintf("/clusters/%d", clusterID)
	resp, err := c.DoRequest(http.MethodPut, path, req)
	if err != nil {
		return nil, err
	}

	var response ClusterUpdateResponse
	if err := c.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response.Cluster, nil
}

// DeleteCluster deletes a cluster
func (c *Client) DeleteCluster(clusterID int) error {
	path := fmt.Sprintf("/clusters/%d", clusterID)
	resp, err := c.DoRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}
