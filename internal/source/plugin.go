package botkubeplugin

import (
	"context"
	"github.com/hashicorp/go-plugin"
	botkubeplugin "github.com/huseyinbabal/bokube-plugins/internal/source/proto"
	"google.golang.org/grpc"
)

type SourcePlugin struct {
	plugin.NetRPCUnsupportedPlugin
	plugin.GRPCPlugin
	Impl Source
}

func (p *SourcePlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	botkubeplugin.RegisterSourceServer(s, &SourceGRPCServer{
		Impl:   p.Impl,
		Broker: broker,
	})
	return nil
}

func (p *SourcePlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &SourceGRPCClient{
		Client: botkubeplugin.NewSourceClient(c),
		Broker: broker,
	}, nil
}

var _ plugin.GRPCPlugin = &SourcePlugin{}
