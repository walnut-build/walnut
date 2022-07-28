package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"

	plugint "github.com/walnut-build/walnut/plugin"
)

// pluginMap is the map of plugins we can dispense.
var pluginMap = map[string]plugin.Plugin {
	"echo": &plugint.TaskPlugin{},
}

func main() {  
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "plugin",
		Output: os.Stdout,
		Level:  hclog.Debug,
	})

	// We're a host! Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: plugint.TaskHandshakeConfig,
		Plugins:         pluginMap,
		Cmd:             exec.Command("./plugin/implementation/echo/echo"),
		Logger:          logger,
	})
	defer client.Kill()

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		log.Fatal(err)
	}

	// Request the plugin
	raw, err := rpcClient.Dispense("echo")
	if err != nil {
		log.Fatal(err)
	}

	// We should have a Greeter now! This feels like a normal interface
	// implementation but is in fact over an RPC connection.
	task := raw.(plugint.Task)
	fmt.Println(task.Run(plugint.RunParameters{
		Cwd: "",
		Arguments: map[string]string {
			"message": "This is some message",
		},
	}))
} 
