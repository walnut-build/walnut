package plugin

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

type RunParameters struct {
	Cwd string
	Arguments map[string]string
}

type TaskResult struct {
	Success bool
	ErrorDescription string
	Artifacts map[string]string
}

type Task interface {
	Run(RunParameters) TaskResult
}

var TaskHandshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "WALNUT_TASK",
	MagicCookieValue: "MTXPTRXXJY",
}

// Here is an implementation that talks over RPC
type TaskRpc struct { 
	client *rpc.Client
}

func (t *TaskRpc) Run(params RunParameters) TaskResult {
	var result TaskResult
	err := t.client.Call("Plugin.Run", params, &result)
	if err != nil {
		panic(err)
	}

	return result
}

type TaskRpcServer struct {
	Impl Task
}

func (s *TaskRpcServer) Run(params RunParameters, result *TaskResult) error {
	*result = s.Impl.Run(params)
	return nil
}

type TaskPlugin struct {
	Impl Task
}

func (p *TaskPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &TaskRpcServer{Impl: p.Impl}, nil
}

func (TaskPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &TaskRpc{client: c}, nil
}
