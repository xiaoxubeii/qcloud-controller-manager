package qcloud

// QCloud is an implementation of cloud provider Interface for Tencent Cloud.
type QCloud struct {
	cvm    CVM
	//clb    CLB
	region string
}

// Initialize provides the cloud with a kubernetes client builder and may spawn goroutines
// to perform housekeeping activities within the cloud provider.
func (c *QCloud) Initialize(clientBuilder controller.ControllerClientBuilder) {}

// LoadBalancer returns a balancer interface. Also returns true if the interface is supported, false otherwise.
func (c *QCloud) LoadBalancer() (LoadBalancer, bool) {
	return c, true
}

// Instances returns an instances interface. Also returns true if the interface is supported, false otherwise.
func (c *QCloud) Instances() (Instances, bool) {
	return c, true
}

// Zones returns a zones interface. Also returns true if the interface is supported, false otherwise.
func (c *QCloud) Zones() (Zones, bool) {
	return c, true
}

// Clusters returns a clusters interface.  Also returns true if the interface is supported, false otherwise.
func (c *QCloud) Clusters() (Clusters, bool) {
	return c, true
}

// Routes returns a routes interface along with whether the interface is supported.
func (c *QCloud) Routes() (Routes, bool) {
	return c, true
}

// ProviderName returns the cloud provider ID.
func (c *QCloud) ProviderName() string {
	return "qcloud"
}

// HasClusterID returns true if a ClusterID is required and set
func (c *QCloud) HasClusterID() bool {
	return true
}
