package backup

import (
	"context"
	"test_plugin/proto"

	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

type BackupGRPCPlugin struct {
	plugin.Plugin
	plugin.GRPCPlugin

	Impl BackupService
}

// GRPCServer is called for plugin server side
func (p *BackupGRPCPlugin) GRPCServer(b *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterBackupServer(s, &BackupServer{Impl: p.Impl})
	return nil
}

// GRPCClient should return the interface implementation for the plugin
// you're serving via gRPC. The provided context will be canceled by
// go-plugin in the event of the plugin process exiting.
func (p *BackupGRPCPlugin) GRPCClient(ctx context.Context, b *plugin.GRPCBroker, cc *grpc.ClientConn) (interface{}, error) {
	return &BackupClient{c: proto.NewBackupClient(cc)}, nil
}
