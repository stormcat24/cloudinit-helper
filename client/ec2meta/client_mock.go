package ec2meta

import (
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"time"
)

type ClientMock struct {
}

func (c ClientMock) GetInstanceIdentityDocument() (*ec2metadata.EC2InstanceIdentityDocument, error) {

	return &ec2metadata.EC2InstanceIdentityDocument{
		DevpayProductCodes: []string{},
		AvailabilityZone: "ap-northeast-1a",
		PrivateIP: "10.0.0.1",
		Version: "2010-08-31",
		Region: "ap-northeast-1",
		InstanceID: "i-fffffff",
		BillingProducts: []string{},
		InstanceType: "t2.nano",
		AccountID: "XXXXXXXXXXXX",
		PendingTime: time.Now(),
		ImageID: "ami-ffffffff",
		KernelID: "",
		RamdiskID: "",
		Architecture: "x86_64",
	}, nil
}