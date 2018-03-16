package qcloud

import (
	"testing"
	"os"
)

func init() {
	os.Setenv("QCloudSecretId", "AKID78zfBWHkeNuGRRzT4Z0Mky1ta6Memb0I")
	os.Setenv("QCloudSecretKey", "Ll0BD4ta6XnrmWJh1CC5tT1NYYOWVQGw")
	os.Setenv("QCloudCvmAPIRegion", "ap-beijing")
}

func TestDescribeInstances(t *testing.T) {
	client := NewCVMFromEnv()
	args := &DescribeInstancesArgs{}
	response, err := client.DescribeInstances(args)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%++v", response.InstanceSet)

}
