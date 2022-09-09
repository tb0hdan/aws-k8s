package external

import (
	"context"
	"os/user"

	"aws-k8s/pkg/utils"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	log "github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
)

type STSClient struct {
	Client     *sts.Client
	User       *user.User
	Role       string
	AssumeRole string
}

func (a *STSClient) Get() *STSClient {
	absolutePath, err := utils.Expand("~/.aws/aws-k8s.ini", a.User)
	if err != nil {
		log.Fatalf("Could not expand path: %+v\n", err)
	}
	iniCfg, err := ini.Load(absolutePath)
	if err != nil {
		log.Fatalf("Could not read default config: %+v\n", err)
	}
	section, err := iniCfg.GetSection("default")
	if err != nil {
		log.Fatalf("Config doesn't have default section: %+v\n", err)
	}

	key, err := section.GetKey("region")
	if err != nil {
		log.Fatalf("Config doesn't have region value: %+v\n", err)
	}
	role, err := section.GetKey("role_arn")
	if err != nil {
		log.Fatalf("Config doesn't have role_arn value: %+v\n", err)
	}

	assumeRole, err := section.GetKey("assume_arn")
	if err != nil {
		log.Fatalf("Config doesn't have assume_arn value: %+v\n", err)
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(key.Value()))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	return &STSClient{
		Client:     sts.NewFromConfig(cfg),
		Role:       role.Value(),
		AssumeRole: assumeRole.Value(),
	}
}
