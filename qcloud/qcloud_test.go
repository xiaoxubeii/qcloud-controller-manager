package qcloud

import (
	"testing"
	"k8s.io/apimachinery/pkg/types"
)

var client *QCloud
var testInstanceName = ""
var testInstanceId = ""

func init() {
	client = FakeNewQCloud()
}

func commonTest(model interface{}, err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%++v", model)
}

func TestNodeAddresses(t *testing.T) {
	nodeAddress, err := client.NodeAddresses(types.NodeName(testInstanceName))
	commonTest(nodeAddress, err, t)
}

func TestNodeAddressesByProviderID(t *testing.T) {
	nodeAddress, err := client.NodeAddressesByProviderID(testInstanceId)
	commonTest(nodeAddress, err, t)
}

func TestExternalID(t *testing.T) {
	id, err := client.ExternalID(types.NodeName(testInstanceName))
	commonTest(id, err, t)
}

func TestInstanceID(t *testing.T) {
	id, err := client.InstanceID(types.NodeName(testInstanceName))
	commonTest(id, err, t)
}

func TestInstanceType(t *testing.T) {
	ty, err := client.InstanceType(types.NodeName(testInstanceName))
	commonTest(ty, err, t)
}

func TestCurrentNodeName(t *testing.T) {
	name, err := client.CurrentNodeName(testInstanceName)
	commonTest(name, err, t)
}

func TestInstanceExistsByProviderID(t *testing.T) {
	exist, err := client.InstanceExistsByProviderID(testInstanceId)
	commonTest(exist, err, t)
}