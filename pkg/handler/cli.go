package handler

import "github.com/tb0hdan/aws-k8s/pkg/commands"

type CLI struct {
	Debug bool `help:"Enable debug mode."`

	Keys       commands.KeysCmd       `cmd:"keys" help:"Manage AWS session keys"`
	AssumeKeys commands.AssumeKeysCmd `cmd:"assume-keys" help:"Manage AWS Assume role session keys"`
	Refresh    commands.RefreshCmd    `cmd:"refresh" help:"Refresh credentials using MFA"`
	Assume     commands.AssumeCmd     `cmd:"assume" help:"Assume credentials using MFA"`
	Wrap       commands.WrapCmd       `cmd:"wrap" help:"Wrap kubectl with exported variables after assume role" passthrough:""`
	Configure  commands.ConfigureCmd  `cmd:"configure" help:"Run interactive configuration"`
}
