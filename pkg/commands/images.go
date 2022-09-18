package commands

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/tb0hdan/aws-k8s/pkg/commands/images"
	"github.com/tb0hdan/aws-k8s/pkg/commands/kubectl"
	"github.com/tb0hdan/aws-k8s/pkg/config"
	"github.com/tb0hdan/aws-k8s/pkg/external"
)

type ImageCmd struct {
	Arguments []string `arg:"" optional:"" name:"arguments" help:"Kubectl arguments"`
}

func (i *ImageCmd) Run(ctx *CLIContext) error {
	var (
		kubeGetPodsResponse images.GetPodsResponse
	)
	r := kubectl.New("~/.aws/aws-k8s-assume.json")
	args := make([]string, 0, len(i.Arguments))
	args = append(args, i.Arguments...)
	args = append(args, []string{"get", "pods", "-o", "json"}...)
	out, err := r.Run(args...)

	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(out), &kubeGetPodsResponse)
	if err != nil {
		return err
	}

	for _, item := range kubeGetPodsResponse.Items {
		fmt.Println(item.Metadata.Name, item.Spec.Containers[0].Image)
	}
	//
	appConfig := config.New("~/.aws/aws-k8s.ini")
	appConfig.Load()
	ecrClient := &external.ECRClient{AppConfig: appConfig}
	client := ecrClient.Get()

	repositories, err := client.Client.DescribeRepositories(context.Background(), &ecr.DescribeRepositoriesInput{
		MaxResults: nil,
		// NextToken:       nil,
		// RegistryId:      nil,
		// RepositoryNames: nil,
	})
	if err != nil {
		return err
	}

	for _, repository := range repositories.Repositories {
		images, err := client.Client.ListImages(context.Background(), &ecr.ListImagesInput{
			RepositoryName: repository.RepositoryName,
			// Filter:         nil,
			// MaxResults:     nil,
			// NextToken:      nil,
			// RegistryId:     nil,
		})
		if err != nil {
			continue
		}
		for _, image := range images.ImageIds {
			fmt.Println(*image.ImageTag, *image.ImageDigest)
		}
	}

	return nil
}
