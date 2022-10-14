package botkubeplugin

import (
	"context"
	"github.com/hashicorp/go-plugin"
	botkubeplugin "github.com/huseyinbabal/botkube-plugins/internal/executor/proto"
)

type ExecutorGRPCServer struct {
	Impl   Executor
	Broker *plugin.GRPCBroker
	botkubeplugin.UnimplementedExecutorServer
}

func (p *ExecutorGRPCServer) Execute(ctx context.Context, request *botkubeplugin.ExecuteRequest) (*botkubeplugin.ExecuteResponse, error) {
	result, err := p.Impl.Execute(request.Command)
	if err != nil {
		return nil, err
	}
	return &botkubeplugin.ExecuteResponse{Data: result}, nil
}

type ExecutorGRPCClient struct {
	Broker *plugin.GRPCBroker
	Client botkubeplugin.ExecutorClient
}

func (p *ExecutorGRPCClient) Execute(command string) (string, error) {
	res, err := p.Client.Execute(context.Background(), &botkubeplugin.ExecuteRequest{Command: command})
	if err != nil {
		return "", err
	}
	return res.Data, nil
}
