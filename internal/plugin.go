package internal

import "github.com/hashicorp/go-plugin"

var handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BOTKUBE_MAGIC_COOKIE",
	MagicCookieValue: "BOTKUBE_BASIC_PLUGIN",
}

func Serve(p map[string]plugin.Plugin) {
	plugin.Serve(&plugin.ServeConfig{
		Plugins:         p,
		HandshakeConfig: handshake,
		GRPCServer:      plugin.DefaultGRPCServer,
	})
}
