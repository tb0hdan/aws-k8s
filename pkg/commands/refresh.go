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

type RefreshCmd struct {
	Token string `cmd:"token" help:"MFA token"`
}

func (r *RefreshCmd) Run(ctx *CLIContext) error {
	appConfig := config.New("~/.aws/aws-k8s.ini")
	appConfig.Load()
	stsClient := &external.STSClient{User: ctx.User, AppConfig: appConfig}
	client := stsClient.Get()
	tokenInfo := &sts.GetSessionTokenInput{
		DurationSeconds: utils.Int32Ptr(3600),
		SerialNumber:    utils.StrPtr(client.Role),
		TokenCode:       utils.StrPtr(r.Token),
	}
	output, err := client.Client.GetSessionToken(context.Background(), tokenInfo)
	if err != nil {
		log.Fatalf("GetSessionToken failed with: %+v\n", err)
	}

	absolutePath := kong.ExpandPath("~/.aws/aws-k8s.json")
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
