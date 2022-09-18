package kubectl

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/alecthomas/kong"
	"github.com/pkg/errors"
	"github.com/tb0hdan/aws-k8s/pkg/auth"
)

type Runner struct {
	CredentialsFile string
}

func (r *Runner) Run(arguments ...string) (output string, err error) {
	var (
		AccessKeyId     string
		SecretAccessKey string
		SessionToken    string
	)

	absolutePath := kong.ExpandPath(r.CredentialsFile)
	credentials := auth.NewCredentials(absolutePath)
	if credentials.Valid() {
		validCredentials, err := credentials.Load()
		if err != nil {
			return "", errors.Wrapf(err, "Could not load credentials")
		}
		AccessKeyId = validCredentials.AccessKeyId
		SecretAccessKey = validCredentials.SecretAccessKey
		SessionToken = validCredentials.SessionToken

	} else {
		return "", fmt.Errorf("credentials invalid, please run `aws-k8s assume --token=123456`")
	}
	os.Setenv("AWS_ACCESS_KEY_ID", AccessKeyId)
	os.Setenv("AWS_SECRET_ACCESS_KEY", SecretAccessKey)
	os.Setenv("AWS_SESSION_TOKEN", SessionToken)
	cmd := exec.Command("kubectl", arguments...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return string(out), err
	}
	return string(out), nil
}

func New(credentialsFile string) *Runner {
	return &Runner{CredentialsFile: credentialsFile}
}
