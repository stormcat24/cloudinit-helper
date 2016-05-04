package ec2

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type Client interface {
	DescribeInstance(instaceId string) (*ec2.Instance, error)
}

type ClientImpl struct {
	service *ec2.EC2
}

func (c ClientImpl) DescribeInstance(instanceId string) (*ec2.Instance, error) {

	input := ec2.DescribeInstancesInput{
		InstanceIds: []*string{
			aws.String(instanceId),
		},
	}

	output, err := c.service.DescribeInstances(&input)
	if err != nil {
		return nil, err
	}

	if len (output.Reservations) > 0 {
		if len (output.Reservations[0].Instances) > 0 {
			return output.Reservations[0].Instances[0], nil
		}
	}

	return nil, fmt.Errorf("Not found specified instance: Id=%v", instanceId)
}