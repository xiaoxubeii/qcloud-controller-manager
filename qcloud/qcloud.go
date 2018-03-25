package qcloud

import (
	"k8s.io/kubernetes/pkg/cloudprovider"
	"k8s.io/kubernetes/pkg/controller"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"fmt"
	"gopkg.in/gcfg.v1"
	"io"
	"os"
)

// QCloud is an implementation of cloud provider Interface for Tencent Cloud.
type QCloud struct {
	cvm CVM
}

type CloudConfig struct {
	Global struct {
		SecretID  string
		SecretKey string
		Region    string
		Host      string
		Path      string
	}
}

const ProviderName = "qcloud"

func init() {
	cloudprovider.RegisterCloudProvider(ProviderName, func(config io.Reader) (cloudprovider.Interface, error) {
		cfg, err := readConfig(config)
		if err != nil {
			return nil, err
		}
		return newQCloud(cfg), nil
	})
}

func readConfig(config io.Reader) (CloudConfig, error) {
	if config == nil {
		return CloudConfig{}, fmt.Errorf("no QCloud provider config file given")
	}
	cfg := configFromEnv()
	err := gcfg.ReadInto(&cfg, config)
	return cfg, err
}

func configFromEnv() (config CloudConfig) {
	config.Global.SecretID = os.Getenv("QCloudSecretId")
	config.Global.SecretKey = os.Getenv("QCloudSecretKey")
	config.Global.Region = os.Getenv("QCloudCvmAPIRegion")
	config.Global.Host = os.Getenv("QCloudCvmAPIHost")
	config.Global.Path = os.Getenv("QCloudCvmAPIPath")

	return
}

func newQCloud(config CloudConfig) *QCloud {
	q := new(QCloud)
	q.cvm = NewCVM(config)
	return q
}

func (c *QCloud) Instances() (cloudprovider.Instances, bool) {
	return c, true
}

func (c *QCloud) Initialize(clientBuilder controller.ControllerClientBuilder) {

}

// LoadBalancer returns a balancer interface. Also returns true if the interface is supported, false otherwise.
func (c *QCloud) LoadBalancer() (cloudprovider.LoadBalancer, bool) {
	return nil, false
}

// Zones returns a zones interface. Also returns true if the interface is supported, false otherwise.
func (c *QCloud) Zones() (cloudprovider.Zones, bool) {
	return nil, false
}

// Clusters returns a clusters interface.  Also returns true if the interface is supported, false otherwise.
func (c *QCloud) Clusters() (cloudprovider.Clusters, bool) {
	return nil, false
}

// Routes returns a routes interface along with whether the interface is supported.
func (c *QCloud) Routes() (cloudprovider.Routes, bool) {
	return nil, false
}

// ProviderName returns the cloud provider ID.
func (c *QCloud) ProviderName() string {
	return ProviderName
}

// ScrubDNS provides an opportunity for cloud-provider-specific code to process DNS settings for pods.
func (c *QCloud) ScrubDNS(nameservers, searches []string) (nsOut, srchOut []string) {
	return
}

// HasClusterID returns true if a ClusterID is required and set
func (c *QCloud) HasClusterID() bool {
	return false
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
