package external

import (
	"context"
	"os/user"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	log "github.com/sirupsen/logrus"
	appConfig "github.com/tb0hdan/aws-k8s/pkg/config"
)

// https://github.com/aws/aws-sdk-go-v2/tree/main/service/sts
type STSClient struct {
	AppConfig  *appConfig.Application
	Client     *sts.Client
	User       *user.User
	Role       string
	AssumeRole string
}

func (a *STSClient) Get() *STSClient {
	region := a.AppConfig.GetRegion()
	if len(region) == 0 {
		log.Fatal("Config doesn't have region value")
	}
	roleARN := a.AppConfig.GetRoleARN()
	if len(roleARN) == 0 {
		log.Fatal("Config doesn't have role_arn value")
	}

	assumeARN := a.AppConfig.GetAssumeARN()
	if len(assumeARN) == 0 {
		log.Fatal("Config doesn't have assume_arn value")
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	return &STSClient{
		Client:     sts.NewFromConfig(cfg),
		Role:       roleARN,
		AssumeRole: assumeARN,
	}
}
