package ec2

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func NewClient(region string) Client {
	// TODO mock

	cred := credentials.NewChainCredentials([]credentials.Provider{
		&credentials.EnvProvider{},
		&credentials.SharedCredentialsProvider{Filename: "", Profile: ""},
		&ec2rolecreds.EC2RoleProvider{ExpiryWindow: 5 * time.Minute},
	})
	conf := aws.NewConfig().WithCredentials(cred).WithMaxRetries(aws.UseServiceDefaultRetries).WithRegion(region)

	return ClientImpl{
		service: ec2.New(session.New(conf)),
	}
}
