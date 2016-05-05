package ec2

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
)

func NewClient(region string) Client {
	// TODO mock
	ses := session.New()
	cred := credentials.NewChainCredentials([]credentials.Provider{
		&credentials.EnvProvider{},
		&ec2rolecreds.EC2RoleProvider{
			Client: ec2metadata.New(ses),
			ExpiryWindow: 5 * time.Minute,
		},
	})
	ses.Config.Credentials = cred
	ses.Config.WithMaxRetries(aws.UseServiceDefaultRetries).WithRegion(region)

	return ClientImpl{
		service: ec2.New(ses),
	}
}
