package qcloud

// QCloud is an implementation of cloud provider Interface for Tencent Cloud.
type QCloud struct {
	cvm CVM
	//clb    CLB
	region string
}

// Initialize provides the cloud with a kubernetes client builder and may spawn goroutines
// to perform housekeeping activities within the cloud provider.
//func (c *QCloud) Initialize(clientBuilder controller.ControllerClientBuilder) {}
//
//// LoadBalancer returns a balancer interface. Also returns true if the interface is supported, false otherwise.
//func (c *QCloud) LoadBalancer() (cloudprovider.LoadBalancer, bool) {
//	return c, true
//}
//
//func (c *QCloud) GetLoadBalancer(clusterName string, service *v1.Service) (status *v1.LoadBalancerStatus, exists bool, err error) {
//	return nil, true, nil
//}
//
//func (c *QCloud) EnsureLoadBalancer(clusterName string, service *v1.Service, nodes []*v1.Node) (*v1.LoadBalancerStatus, error) {
//	return nil, nil
//}
//
//func (c *QCloud) UpdateLoadBalancer(clusterName string, service *v1.Service, nodes []*v1.Node) error {
//	return nil
//}
//
//func (c *QCloud) EnsureLoadBalancerDeleted(clusterName string, service *v1.Service) error {
//	return nil
//}

//// Instances returns an instances interface. Also returns true if the interface is supported, false otherwise.
//func (c *QCloud) Instances() (Instances, bool) {
//	return c, true
//}
//
//// Zones returns a zones interface. Also returns true if the interface is supported, false otherwise.
//func (c *QCloud) Zones() (Zones, bool) {
//	return c, true
//}
//
//// Clusters returns a clusters interface.  Also returns true if the interface is supported, false otherwise.
//func (c *QCloud) Clusters() (Clusters, bool) {
//	return c, true
//}
//
//// Routes returns a routes interface along with whether the interface is supported.
//func (c *QCloud) Routes() (Routes, bool) {
//	return c, true
//}
//
//// ProviderName returns the cloud provider ID.
//func (c *QCloud) ProviderName() string {
//	return "qcloud"
//}

// HasClusterID returns true if a ClusterID is required and set
//func (c *QCloud) HasClusterID() bool {
//	return true
//}

//func (c *QCloud) NodeAddresses(name types.NodeName) ([]v1.NodeAddress, error) {
//	return nil, nil
//}

//func (c *QCloud) NodeAddressesByProviderID(providerID string) ([]v1.NodeAddress, error)
//func (c *QCloud) ExternalID(nodeName types.NodeName) (string, error)
//func (c *QCloud) InstanceID(nodeName types.NodeName) (string, error)
//func (c *QCloud) InstanceType(name types.NodeName) (string, error)
//func (c *QCloud) InstanceTypeByProviderID(providerID string) (string, error)
//func (c *QCloud) AddSSHKeyToAllInstances(user string, keyData []byte) error
//func CurrentNodeName(hostname string) (types.NodeName, error)
//func InstanceExistsByProviderID(providerID string) (bool, error)
