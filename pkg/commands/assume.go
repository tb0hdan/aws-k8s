package commands

import (
	"context"

	"github.com/alecthomas/kong"
	"github.com/tb0hdan/aws-k8s/pkg/auth"
	"github.com/tb0hdan/aws-k8s/pkg/config"
	"github.com/tb0hdan/aws-k8s/pkg/external"
	"github.com/tb0hdan/aws-k8s/pkg/utils"

	"github.com/aws/aws-sdk-go-v2/service/sts"
	log "github.com/sirupsen/logrus"
)

type AssumeCmd struct {
	Token string `cmd:"token" help:"MFA token"`
}

func (a *AssumeCmd) Run(ctx *CLIContext) error {
	appConfig := config.New("~/.aws/aws-k8s.ini")
	appConfig.Load()
	stsClient := &external.STSClient{User: ctx.User, AppConfig: appConfig}
	client := stsClient.Get()
	output, err := client.Client.AssumeRole(context.Background(), &sts.AssumeRoleInput{
		RoleArn:         utils.StrPtr(client.AssumeRole),
		RoleSessionName: utils.StrPtr("AWSCLI-Session"),
		DurationSeconds: utils.Int32Ptr(3600),
		SerialNumber:    utils.StrPtr(client.Role),
		TokenCode:       utils.StrPtr(a.Token),
	})
	if err != nil {
		log.Fatalf("AssumeRole failed with: %+v\n", err)
	}

	absolutePath := kong.ExpandPath("~/.aws/aws-k8s-assume.json")
	credentials := auth.NewCredentials(absolutePath)
	err = credentials.Save(&auth.CredentialsCache{
		AccessKeyId:     *output.Credentials.AccessKeyId,
		Expiration:      *output.Credentials.Expiration,
		SecretAccessKey: *output.Credentials.SecretAccessKey,
		SessionToken:    *output.Credentials.SessionToken,
	})
	if err != nil {
		log.Fatalf("Could not save credentials to cache: %+v\n", err)
	}
	return nil
}
