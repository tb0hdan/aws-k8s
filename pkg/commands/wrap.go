package commands

import (
	"fmt"

	"github.com/tb0hdan/aws-k8s/pkg/commands/kubectl"
)

type WrapCmd struct {
	Arguments []string `arg:"" optional:"" name:"arguments" help:"Kubectl arguments"`
}

func (w *WrapCmd) Run(ctx *CLIContext) error {
	r := kubectl.New("~/.aws/aws-k8s-assume.json")
	out, err := r.Run(w.Arguments...)
	if err != nil {
		return err
	}
	fmt.Println(out)
	return nil
}
