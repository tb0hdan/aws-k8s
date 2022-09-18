package external

import (
	"github.com/alecthomas/kong"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/tb0hdan/aws-k8s/pkg/auth"
	appConfig "github.com/tb0hdan/aws-k8s/pkg/config"
)

// https://github.com/aws/aws-sdk-go-v2/tree/main/service/ecr
type ECRClient struct {
	AppConfig *appConfig.Application
	Client    *ecr.Client
}

func (e *ECRClient) Get() *ECRClient {
	absolutePath := kong.ExpandPath("~/.aws/aws-k8s.json")
	cache := auth.NewCredentials(absolutePath)
	loaded, err := cache.Load()
	if err != nil {
		panic(err)
	}

	client := ecr.New(ecr.Options{
		Credentials: loaded,
		Region:      e.AppConfig.GetRegion(),
	})
	return &ECRClient{Client: client}
}
