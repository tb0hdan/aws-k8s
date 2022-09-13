package commands

import (
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/pkg/errors"
	"github.com/tb0hdan/aws-k8s/pkg/auth"
)

type AssumeKeysCmd struct {
	Print bool `cmd:"print" help:"Print AWS values"`
}

func (p *AssumeKeysCmd) Run(ctx *CLIContext) error {
	var (
		AccessKeyId     string
		SecretAccessKey string
		SessionToken    string
	)
	if !p.Print {
		return nil
	}

	absolutePath := kong.ExpandPath("~/.aws/aws-k8s-assume.json")
	credentials := auth.NewCredentials(absolutePath)
	if credentials.Valid() {
		validCredentials, err := credentials.Load()
		if err != nil {
			return errors.Wrapf(err, "Could not load credentials")
		}
		AccessKeyId = validCredentials.AccessKeyId
		SecretAccessKey = validCredentials.SecretAccessKey
		SessionToken = validCredentials.SessionToken

	} else {
		return fmt.Errorf("credentials invalid, please run `aws-k8s assume --token=123456`")
	}
	fmt.Printf("export AWS_ACCESS_KEY_ID=%s; export AWS_SECRET_ACCESS_KEY=%s; export AWS_SESSION_TOKEN=%s",
		AccessKeyId,
		SecretAccessKey,
		SessionToken,
	)
	return nil
}
