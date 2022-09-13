package external

import (
	"github.com/alecthomas/kong"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/tb0hdan/aws-k8s/pkg/auth"
	appConfig "github.com/tb0hdan/aws-k8s/pkg/config"
)

type EKSClient struct {
	AppConfig *appConfig.Application
	Client    *eks.Client
}

func (e *EKSClient) Get() *EKSClient {
	absolutePath := kong.ExpandPath("~/.aws/aws-k8s.json")
	cache := auth.NewCredentials(absolutePath)
	loaded, err := cache.Load()
	if err != nil {
		panic(err)
	}

	client := eks.New(eks.Options{
		Credentials: loaded,
		Region:      e.AppConfig.GetRegion(),
	})
	return &EKSClient{Client: client}
}
