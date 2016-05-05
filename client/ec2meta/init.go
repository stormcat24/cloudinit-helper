package ec2meta

import (
	"time"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
)

type Config struct {
	UseMock bool
}

func NewClient(mock bool) Client {

	if mock {
		return ClientMock{}
	} else {
		ses := session.New()
		cred := credentials.NewChainCredentials([]credentials.Provider{
			&credentials.EnvProvider{},
			&ec2rolecreds.EC2RoleProvider{
				Client: ec2metadata.New(ses),
				ExpiryWindow: 5 * time.Minute,
			},
		})
		ses.Config.Credentials = cred
		ses.Config.WithMaxRetries(aws.UseServiceDefaultRetries)

		return ClientImpl{
			service: ec2metadata.New(ses),
		}
	}
}

