package commands

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/tb0hdan/aws-k8s/pkg/auth"
	"github.com/tb0hdan/aws-k8s/pkg/utils"

	"github.com/pkg/errors"
)

type WrapCmd struct {
	Arguments []string `arg:"" optional:"" name:"arguments" help:"Kubectl arguments"`
}

func (w *WrapCmd) Run(ctx *CLIContext) error {
	var (
		AccessKeyId     string
		SecretAccessKey string
		SessionToken    string
	)

	absolutePath, err := utils.Expand("~/.aws/aws-k8s-assume.json", ctx.User)
	if err != nil {
		return errors.Wrapf(err, "Could not expand path")
	}
	credentials := auth.NewCredentials(absolutePath)
	if credentials.Valid() {
		validCredentials, err := credentials.Load()
		if err != nil {
			return errors.Wrapf(err, "Could not load credentials")
		}
		AccessKeyId = validCredentials.AccessKeyId
		SecretAccessKey = validCredentials.SecretAccessKey
		SessionToken = validCredentials.SessionToken

	} else {
		return fmt.Errorf("credentials invalid, please run `aws-k8s assume --token=123456`")
	}
	os.Setenv("AWS_ACCESS_KEY_ID", AccessKeyId)
	os.Setenv("AWS_SECRET_ACCESS_KEY", SecretAccessKey)
	os.Setenv("AWS_SESSION_TOKEN", SessionToken)
	cmd := exec.Command("kubectl", w.Arguments...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println(string(out))
	return nil
}
