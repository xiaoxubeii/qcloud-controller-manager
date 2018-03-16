package qcloud

import (
	"github.com/dbdd4us/qcloudapi-sdk-go/common"
	"github.com/dbdd4us/qcloudapi-sdk-go/cvm"
	"os"
)

func NewCVMFromEnv() *CVM {
	secretId := os.Getenv("QCloudSecretId")
	secretKey := os.Getenv("QCloudSecretKey")
	region := os.Getenv("QCloudCvmAPIRegion")
	return commonCred(secretId, secretKey, region)
}

func commonCred(secret_id, secretKey, region string) *CVM {
	credential := common.Credential{
		SecretId:  secret_id,
		SecretKey: secretKey,
	}

	opts := common.Opts{
		Region: region,
	}
	client, _ := cvm.NewClient(credential, opts)
	return &CVM{client}
}

type CVM struct {
	*cvm.Client
}

type DescribeInstancesArgs struct {
	*cvm.DescribeInstancesArgs
}

type DescribeInstancesResponse struct {
	*cvm.DescribeInstancesResponse
}

func (c *CVM) DescribeInstances(args *DescribeInstancesArgs) (*DescribeInstancesResponse, error) {
	response, error := c.Client.DescribeInstances(args.DescribeInstancesArgs)
	return &DescribeInstancesResponse{response}, error
}
