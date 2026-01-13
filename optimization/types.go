package optimization

// Trial represents a trial configuration
type Trial struct {
	ID        int    `json:"id"`
	ProjectID string `json:"project_id"`
	Key       string `json:"key_"`
	Value     string `json:"value_"`
}

// Cluster represents a cluster
type Cluster struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	CreatedAt   string `json:"created_at"`           // String olarak değiştir
	UpdatedAt   string `json:"updated_at,omitempty"` // String olarak değiştir
}

// ClusterCreate represents cluster creation request
type ClusterCreate struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// ClusterUpdate represents cluster update request
type ClusterUpdate struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// ClusterHost represents a host in a cluster
type ClusterHost struct {
	ID        string `json:"id"`
	ClusterID string `json:"cluster_id"`
	Hostname  string `json:"hostname"`
	Enabled   bool   `json:"enabled"`
	CreatedAt string `json:"created_at"`           // String olarak değiştir
	UpdatedAt string `json:"updated_at,omitempty"` // String olarak değiştir
}

// ClusterHostCreate represents host creation request
type ClusterHostCreate struct {
	Hostname string `json:"hostname"`
	Enabled  bool   `json:"enabled"`
}

// ClusterHostUpdate represents host update request
type ClusterHostUpdate struct {
	Hostname string `json:"hostname,omitempty"`
	Enabled  *bool  `json:"enabled,omitempty"`
}

// ClusterExcludedVM represents an excluded VM
type ClusterExcludedVM struct {
	ID        string `json:"id"`
	ClusterID string `json:"cluster_id"`
	VMName    string `json:"vm_name"`
	CreatedAt string `json:"created_at"`           // String olarak değiştir
	UpdatedAt string `json:"updated_at,omitempty"` // String olarak değiştir
}

// ClusterExcludedVMCreate represents excluded VM creation request
type ClusterExcludedVMCreate struct {
	VMName string `json:"vm_name"`
}

// HostMaintenancePolicy represents a host maintenance policy
type HostMaintenancePolicy struct {
	ID        string `json:"id"`
	ClusterID string `json:"cluster_id"`
	Name      string `json:"name"`
	Enabled   bool   `json:"enabled"`
	CreatedAt string `json:"created_at"`           // String olarak değiştir
	UpdatedAt string `json:"updated_at,omitempty"` // String olarak değiştir
}

// HostMaintenancePolicyCreate represents policy creation request
type HostMaintenancePolicyCreate struct {
	ClusterID string `json:"cluster_id"`
	Name      string `json:"name"`
	Enabled   bool   `json:"enabled"`
}

// HostMaintenancePolicyUpdate represents policy update request
type HostMaintenancePolicyUpdate struct {
	ClusterID string `json:"cluster_id,omitempty"`
	Name      string `json:"name,omitempty"`
	Enabled   *bool  `json:"enabled,omitempty"`
}

// WorkloadBalancingPolicy represents a workload balancing policy
type WorkloadBalancingPolicy struct {
	ID              string `json:"id"`
	ClusterID       string `json:"cluster_id"`
	Name            string `json:"name"`
	BalancingMode   string `json:"balancing_mode"`
	CPUBalancing    bool   `json:"cpu_balancing"`
	MemoryBalancing bool   `json:"memory_balancing"`
	Period          int    `json:"period"`
	Enabled         bool   `json:"enabled"`
	CreatedAt       string `json:"created_at"`           // String olarak değiştir
	UpdatedAt       string `json:"updated_at,omitempty"` // String olarak değiştir
}

// WorkloadBalancingPolicyCreate represents policy creation request
type WorkloadBalancingPolicyCreate struct {
	ClusterID       string `json:"cluster_id"`
	Name            string `json:"name"`
	BalancingMode   string `json:"balancing_mode"`
	CPUBalancing    bool   `json:"cpu_balancing"`
	MemoryBalancing bool   `json:"memory_balancing"`
	Period          int    `json:"period"`
	Enabled         bool   `json:"enabled"`
}

// WorkloadBalancingPolicyUpdate represents policy update request
type WorkloadBalancingPolicyUpdate struct {
	ClusterID       string `json:"cluster_id,omitempty"`
	Name            string `json:"name,omitempty"`
	BalancingMode   string `json:"balancing_mode,omitempty"`
	CPUBalancing    *bool  `json:"cpu_balancing,omitempty"`
	MemoryBalancing *bool  `json:"memory_balancing,omitempty"`
	Period          int    `json:"period,omitempty"`
	Enabled         *bool  `json:"enabled,omitempty"`
}

// WorkloadConsolidationPolicy represents a workload consolidation policy
type WorkloadConsolidationPolicy struct {
	ID        string `json:"id"`
	ClusterID string `json:"cluster_id"`
	Name      string `json:"name"`
	Period    int    `json:"period"`
	Enabled   bool   `json:"enabled"`
	CreatedAt string `json:"created_at"`           // String olarak değiştir
	UpdatedAt string `json:"updated_at,omitempty"` // String olarak değiştir
}

// WorkloadConsolidationPolicyCreate represents policy creation request
type WorkloadConsolidationPolicyCreate struct {
	ClusterID string `json:"cluster_id"`
	Name      string `json:"name"`
	Period    int    `json:"period"`
	Enabled   bool   `json:"enabled"`
}

// WorkloadConsolidationPolicyUpdate represents policy update request
type WorkloadConsolidationPolicyUpdate struct {
	ClusterID string `json:"cluster_id,omitempty"`
	Name      string `json:"name,omitempty"`
	Period    int    `json:"period,omitempty"`
	Enabled   *bool  `json:"enabled,omitempty"`
}

// Response wrappers for API responses
type ClusterCreateResponse struct {
	Message string  `json:"message"`
	Code    int     `json:"code"`
	Title   string  `json:"title"`
	Cluster Cluster `json:"cluster"`
}

type ClusterUpdateResponse struct {
	Message string  `json:"message"`
	Code    int     `json:"code"`
	Title   string  `json:"title"`
	Cluster Cluster `json:"cluster"`
}

type ClusterHostCreateResponse struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Title   string      `json:"title"`
	Host    ClusterHost `json:"host"`
}

type ClusterHostUpdateResponse struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Title   string      `json:"title"`
	Host    ClusterHost `json:"host"`
}

type ClusterExcludedVMCreateResponse struct {
	Message string            `json:"message"`
	Code    int               `json:"code"`
	Title   string            `json:"title"`
	VM      ClusterExcludedVM `json:"vm"`
}

type HostMaintenancePolicyCreateResponse struct {
	Message string                `json:"message"`
	Code    int                   `json:"code"`
	Title   string                `json:"title"`
	Policy  HostMaintenancePolicy `json:"policy"`
}

type HostMaintenancePolicyUpdateResponse struct {
	Message string                `json:"message"`
	Code    int                   `json:"code"`
	Title   string                `json:"title"`
	Policy  HostMaintenancePolicy `json:"policy"`
}

type WorkloadBalancingPolicyCreateResponse struct {
	Message string                  `json:"message"`
	Code    int                     `json:"code"`
	Title   string                  `json:"title"`
	Policy  WorkloadBalancingPolicy `json:"policy"`
}

type WorkloadBalancingPolicyUpdateResponse struct {
	Message string                  `json:"message"`
	Code    int                     `json:"code"`
	Title   string                  `json:"title"`
	Policy  WorkloadBalancingPolicy `json:"policy"`
}

type WorkloadConsolidationPolicyCreateResponse struct {
	Message string                      `json:"message"`
	Code    int                         `json:"code"`
	Title   string                      `json:"title"`
	Policy  WorkloadConsolidationPolicy `json:"policy"`
}

type WorkloadConsolidationPolicyUpdateResponse struct {
	Message string                      `json:"message"`
	Code    int                         `json:"code"`
	Title   string                      `json:"title"`
	Policy  WorkloadConsolidationPolicy `json:"policy"`
}
