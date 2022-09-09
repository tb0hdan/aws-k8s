package main

import (
	"os/user"

	"github.com/tb0hdan/aws-k8s/pkg/commands"
	"github.com/tb0hdan/aws-k8s/pkg/utils"

	"github.com/alecthomas/kong"
)

var CLI struct {
	Debug bool `help:"Enable debug mode."`

	Keys       commands.KeysCmd       `cmd:"keys" help:"Manage AWS session keys"`
	AssumeKeys commands.AssumeKeysCmd `cmd:"assume-keys" help:"Manage AWS Assume role session keys"`
	Refresh    commands.RefreshCmd    `cmd:"refresh" help:"Refresh credentials using MFA"`
	Assume     commands.AssumeCmd     `cmd:"assume" help:"Assume credentials using MFA"`
	Wrap       commands.WrapCmd       `cmd:"wrap" help:"Wrap kubectl with exported variables after assume role" passthrough:""`
	Configure  commands.ConfigureCmd  `cmd:"configure" help:"Run interactive configuration"`
}

func main() {
	// clear environment first
	utils.UnsetEnvList("AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY", "AWS_SESSION_TOKEN")
	// get current user
	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}
	//
	cli := kong.Parse(&CLI)
	cli.FatalIfErrorf(cli.Run(&commands.CLIContext{Debug: CLI.Debug, User: currentUser}))
}
