package commands

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/tb0hdan/aws-k8s/pkg/config"
	"github.com/tb0hdan/aws-k8s/pkg/external"
)

type ClustersCmd struct {
}

func (c *ClustersCmd) Run(ctx *CLIContext) error {
	appConfig := config.New("~/.aws/aws-k8s.ini")
	appConfig.Load()

	eksClient := external.EKSClient{
		AppConfig: appConfig,
	}
	client := eksClient.Get()
	clusters, err := client.Client.ListClusters(context.Background(), &eks.ListClustersInput{})
	if err != nil {
		return err
	}

	for _, clusterName := range clusters.Clusters {
		fmt.Println(clusterName)
	}

	return nil
}
