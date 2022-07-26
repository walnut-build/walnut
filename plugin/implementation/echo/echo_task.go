package main

import (
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"

	task "github.com/walnut-build/walnut/plugin/plugint"
)

// Here is a real implementation of Greeter
type GreeterHello struct {
	logger hclog.Logger
}

func (g *GreeterHello) Greet() string {
	g.logger.Debug("message from GreeterHello.Greet")
	return "Hello!"
}

type EchoTask struct {}

func (t *EchoTask) Run(params task.RunParameters) task.TaskResult {
	message := params.Arguments["message"]

	return task.TaskResult {
		Artifacts: map[string]string {
			"message": message,
		},
	}
}

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Trace,
		Output:     os.Stderr,
		JSONFormat: true,
	})

	echoTask := &EchoTask{}

	var pluginMap = map[string]plugin.Plugin {
		"echo": &task.TaskPlugin {
			Impl: echoTask,
		},
	}

	logger.Debug("message from plugin", "foo", "bar")

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: task.TaskHandshakeConfig,
		Plugins:         pluginMap,
	})
}
