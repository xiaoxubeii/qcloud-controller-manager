package qcloud

var fakeInstanceInfo *InstanceInfo

func init() {
	fakeInstanceInfo = NewInstanceInfo()
	fakeInstanceInfo.InstanceID = "ins-9zdg0x3b"
	fakeInstanceInfo.InstanceType = "S2.SMALL2"
	fakeInstanceInfo.InstanceName = "bj-qc-oam-testnewlb-001.tendcloud.com"
	fakeInstanceInfo.PrivateIPAddresses = []string{"172.29.1.227"}
	fakeInstanceInfo.PublicIPAddresses = []string{"172.29.1.227"}
}

func FakeNewCVM() *FakeCVM {
	return &FakeCVM{}
}

type FakeCVM struct {
}

func FakeNewQCloud() *QCloud {
	q := new(QCloud)
	q.cvm = FakeNewCVM()
	return q
}

func (c *FakeCVM) GetInstanceByName(name string) (*InstanceInfo, error) {
	return fakeInstanceInfo, nil
}

func (c *FakeCVM) GetInstanceByID(id string) (*InstanceInfo, error) {
	return fakeInstanceInfo, nil
}
