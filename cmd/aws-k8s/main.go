package main

import "github.com/tb0hdan/aws-k8s/pkg/handler"

func main() {
	cli := &handler.CLI{}
	commandHandler := handler.New()
	commandHandler.ClearEnvironment()
	ctx := commandHandler.Parse(cli)
	commandHandler.Run(ctx, cli.Debug)
}
