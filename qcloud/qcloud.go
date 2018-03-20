package qcloud

import (
	"k8s.io/kubernetes/pkg/cloudprovider"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
)

// QCloud is an implementation of cloud provider Interface for Tencent Cloud.
type QCloud struct {
	cvm *CVM
	//clb    CLB
	region string
}

func NewQCloud() *QCloud {
	q := new(QCloud)
	q.cvm = NewCVMFromEnv()
	return qcloud
}

func (c *QCloud) Instances() (cloudprovider.Instances, bool) {
	return c, true
}

func (c *QCloud) NodeAddresses(name types.NodeName) ([]v1.NodeAddress, error) {
	return nil, nil
}

func (c *QCloud) NodeAddressesByProviderID(providerID string) ([]v1.NodeAddress, error) {
	return nil, nil
}
func (c *QCloud) ExternalID(nodeName types.NodeName) (string, error) {
	return "", nil
}

func (c *QCloud) InstanceID(nodeName types.NodeName) (string, error) {
	return "", nil
}
func (c *QCloud) InstanceType(name types.NodeName) (string, error) {
	return "", nil
}
func (c *QCloud) InstanceTypeByProviderID(providerID string) (string, error) {
	return "", nil
}
func (c *QCloud) AddSSHKeyToAllInstances(user string, keyData []byte) error {
	return nil
}
func (c *QCloud) CurrentNodeName(hostname string) (types.NodeName, error) {
	return "", nil
}
func (c *QCloud) InstanceExistsByProviderID(providerID string) (bool, error) {
	return false, nil
}
