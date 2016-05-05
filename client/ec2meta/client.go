package ec2meta

import (
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
)

type Client interface {
	GetInstanceIdentityDocument() (*ec2metadata.EC2InstanceIdentityDocument, error)
}

type ClientImpl struct {
	service *ec2metadata.EC2Metadata
}

func (c ClientImpl) GetInstanceIdentityDocument() (*ec2metadata.EC2InstanceIdentityDocument, error) {

	output, err := c.service.GetInstanceIdentityDocument()
	if err != nil {
		return nil, err
	}

	return &output, err
}