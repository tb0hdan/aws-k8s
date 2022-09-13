package commands

import (
	"fmt"

	"github.com/alecthomas/kong"
	"gopkg.in/ini.v1"
)

type ConfigureCmd struct {
}

func (c *ConfigureCmd) Run(ctx *CLIContext) error {
	var (
		Region    string
		RoleARN   string
		AssumeARN string
	)
	absolutePath := kong.ExpandPath("~/.aws/aws-k8s.ini")
	//
	fmt.Println("AWS K8S Configuration")
	fmt.Println("Please type required parameters and press enter")
	fmt.Printf("[Region]: ")
	fmt.Scanln(&Region)
	fmt.Printf("[Role ARN]: ")
	fmt.Scanln(&RoleARN)
	fmt.Printf("[Assume ARN]: ")
	fmt.Scanln(&AssumeARN)
	//
	cfg := ini.Empty()
	cfg.Section("default").Key("region").SetValue(Region)
	cfg.Section("default").Key("role_arn").SetValue(RoleARN)
	cfg.Section("default").Key("assume_arn").SetValue(AssumeARN)
	//
	return cfg.SaveTo(absolutePath)
}
