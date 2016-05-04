package meta

type ClientMock struct {
}

func (m ClientMock) GetAvailabilityZone() (*AvailabilityZone, error) {
	return &AvailabilityZone{
		Name: "ap-northeast-1a",
	}, nil
}

func (m ClientMock) GetInstanceID() (string, error) {
	return "i-fffffff", nil
}