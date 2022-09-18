package handler

import "github.com/tb0hdan/aws-k8s/pkg/commands"

type CLI struct {
	// Application wide debug flag
	Debug bool `help:"Enable debug mode."`
	// Actual commands
	Assume     commands.AssumeCmd     `cmd:"assume" help:"Assume credentials using MFA"`
	AssumeKeys commands.AssumeKeysCmd `cmd:"assume-keys" help:"Manage AWS Assume role session keys"`
	Clusters   commands.ClustersCmd   `cmd:"clusters" help:"List EKS clusters"`
	Configure  commands.ConfigureCmd  `cmd:"configure" help:"Run interactive configuration"`
	Keys       commands.KeysCmd       `cmd:"keys" help:"Manage AWS session keys"`
	Refresh    commands.RefreshCmd    `cmd:"refresh" help:"Refresh credentials using MFA"`
	Wrap       commands.WrapCmd       `cmd:"wrap" help:"Wrap kubectl with exported variables after assume role" passthrough:""`
	Images     commands.ImageCmd      `cmd:"images" help:"Get images using kubectl with exported variables after assume role" passthrough:""`
}
