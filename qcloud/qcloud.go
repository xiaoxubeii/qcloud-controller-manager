package qcloud

import (
	"k8s.io/kubernetes/pkg/cloudprovider"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
)

// QCloud is an implementation of cloud provider Interface for Tencent Cloud.
type QCloud struct {
	cvm CVM
}

func newQCloud() *QCloud {
	q := new(QCloud)
	q.cvm = NewCVMFromEnv()
	return q
}

func (c *QCloud) Instances() (cloudprovider.Instances, bool) {
	return c, true
}

func extractInstanceAddress(ins *InstanceInfo, err error) ([]v1.NodeAddress, error) {
	if err != nil {
		return [] v1.NodeAddress{}, err
	}

	var nodeAdds [] v1.NodeAddress

	for _, ip := range ins.PrivateIPAddresses {
		nodeAdds = append(nodeAdds, v1.NodeAddress{Type: v1.NodeInternalIP, Address: ip})
	}

	for _, ip := range ins.PublicIPAddresses {
		nodeAdds = append(nodeAdds, v1.NodeAddress{Type: v1.NodeExternalIP, Address: ip})
	}

	nodeAdds = append(nodeAdds, v1.NodeAddress{Type: v1.NodeHostName, Address: string(ins.InstanceName)})

	return nodeAdds, nil
}

func (c *QCloud) NodeAddresses(name types.NodeName) ([]v1.NodeAddress, error) {
	ins, err := c.cvm.GetInstanceByName(string(name))
	return extractInstanceAddress(ins, err)
}

func (c *QCloud) NodeAddressesByProviderID(providerID string) ([]v1.NodeAddress, error) {
	ins, err := c.cvm.GetInstanceByID(providerID)
	return extractInstanceAddress(ins, err)
}

func (c *QCloud) ExternalID(nodeName types.NodeName) (string, error) {
	ins, err := c.cvm.GetInstanceByName(string(nodeName))
	if err != nil || *ins == (InstanceInfo{}) {
		return "", cloudprovider.InstanceNotFound
	}

	return ins.InstanceID, nil
}

func (c *QCloud) InstanceID(nodeName types.NodeName) (string, error) {
	ins, err := c.cvm.GetInstanceByName(string(nodeName))
	if err != nil || *ins == (InstanceInfo{}) {
		return "", err
	}

	return ins.InstanceID, nil
}

func (c *QCloud) InstanceType(name types.NodeName) (string, error) {
	ins, err := c.cvm.GetInstanceByName(string(name))
	if err != nil {
		return "", err
	}

	return ins.InstanceType, nil
}

func (c *QCloud) InstanceTypeByProviderID(providerID string) (string, error) {
	ins, err := c.cvm.GetInstanceByID(providerID)
	if err != nil {
		return "", err
	}

	return ins.InstanceType, nil
}
func (c *QCloud) AddSSHKeyToAllInstances(user string, keyData []byte) error {
	return cloudprovider.NotImplemented
}
func (c *QCloud) CurrentNodeName(hostname string) (types.NodeName, error) {
	return types.NodeName(hostname), nil
}
func (c *QCloud) InstanceExistsByProviderID(providerID string) (bool, error) {
	ins, err := c.cvm.GetInstanceByID(providerID)

	if err != nil {
		return false, err
	}

	return *ins != (InstanceInfo{}), nil
}
