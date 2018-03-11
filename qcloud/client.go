package qcloud

import (
	"github.com/xiaoxubeii/qcloudapi-sdk-go/common"
)

const SECRET_ID = ""
const SECRET_KEY = ""
const REGION = ""

func NewCVM() CVM {
	credential := common.Credential{
		SecretId:  SECRET_ID,
		SecretKey: SECRET_KEY,
	}

	opts := common.Opts{
		Region: REGION,
	}

	client, _ := clb.NewClient(credential, opts)
	return CVM{client: client}
}

type CVM struct {
	client common.Credential
}

func (c *CVM) DescribeInstances() {

}
