package handler

import (
	"os"
	"os/user"

	"github.com/alecthomas/kong"
	"github.com/tb0hdan/aws-k8s/pkg/commands"
	"github.com/tb0hdan/aws-k8s/pkg/utils"
)

type AWSCommand struct {
	Commands []kong.Option
}

func (a *AWSCommand) ClearEnvironment() {
	// clear environment first
	utils.UnsetEnvList("AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY", "AWS_SESSION_TOKEN")
}

func (a *AWSCommand) AddCommand(name, help, group string, cmd interface{}, tags ...string) {
	a.Commands = append(a.Commands, kong.DynamicCommand(name, help, group, cmd, tags...))
}

func (a *AWSCommand) Parse(cliStruct interface{}) *kong.Context {
	// Show help by default
	if len(os.Args) == 1 {
		os.Args = append(os.Args, "-h")
	}
	return kong.Parse(cliStruct, a.Commands...)
}

func (a *AWSCommand) Run(cli *kong.Context, debug bool) {
	// get current user
	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}
	//
	cli.FatalIfErrorf(cli.Run(&commands.CLIContext{Debug: debug, User: currentUser}))
}

func New() *AWSCommand {
	return &AWSCommand{}
}
