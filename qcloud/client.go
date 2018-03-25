package qcloud

import (
	"github.com/dbdd4us/qcloudapi-sdk-go/common"
	"github.com/dbdd4us/qcloudapi-sdk-go/cvm"
	"os"
)

func NewCVMFromEnv() *CVMImpl {
	secretId := os.Getenv("QCloudSecretId")
	secretKey := os.Getenv("QCloudSecretKey")
	region := os.Getenv("QCloudCvmAPIRegion")
	return commonCred(secretId, secretKey, region)
}

func NewCVM(config CloudConfig) *CVMImpl {
	return commonCred(config.Global.SecretID, config.Global.SecretKey, config.Global.Region)
}

func commonCred(secretId, secretKey, region string) *CVMImpl {
	credential := common.Credential{
		SecretId:  secretId,
		SecretKey: secretKey,
	}

	opts := common.Opts{
		Region: region,
	}
	client, _ := cvm.NewClient(credential, opts)
	return &CVMImpl{client}
}

type CVM interface {
	GetInstanceByName(string) (*InstanceInfo, error)
	GetInstanceByID(string) (*InstanceInfo, error)
}

type CVMImpl struct {
	*cvm.Client
}

type DescribeInstancesArgs struct {
	*cvm.DescribeInstancesArgs
}

type DescribeInstancesResponse struct {
	*cvm.DescribeInstancesResponse
}

type InstanceInfo struct {
	*cvm.InstanceInfo
}

func NewInstanceInfo() *InstanceInfo {
	return &InstanceInfo{&cvm.InstanceInfo{}}
}

func (c *CVMImpl) GetInstanceByName(name string) (*InstanceInfo, error) {
	filters := [] cvm.Filter{cvm.NewFilter(cvm.FilterNameInstanceName, name)}
	version := "2017-03-12"
	args := &cvm.DescribeInstancesArgs{Filters: &filters, Version: version}
	resp, err := c.Client.DescribeInstances(args)
	instances := resp.InstanceSet

	if err != nil || len(resp.InstanceSet) == 0 {
		return &InstanceInfo{}, err
	}

	return &InstanceInfo{&instances[0]}, nil
}

func (c *CVMImpl) GetInstanceByID(id string) (*InstanceInfo, error) {
	version := "2017-03-12"
	args := &cvm.DescribeInstancesArgs{InstanceIds: &[]string{id}, Version: version}
	resp, err := c.Client.DescribeInstances(args)
	instances := resp.InstanceSet

	if err != nil || len(resp.InstanceSet) == 0 {
		return &InstanceInfo{}, err
	}

	return &InstanceInfo{&instances[0]}, nil
}
