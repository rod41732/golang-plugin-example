package backup

import (
	"context"
	"test_plugin/proto"
)

// BackupServer handle receiving request the forward it to `Impl` inside
// It's created when the server process run
type BackupServer struct {
	Impl BackupService
}

var _ proto.BackupServer = (*BackupServer)(nil)

// OnMessageIn delegate logic to `Impl`
func (s *BackupServer) OnMessageIn(ctx context.Context, req *proto.Chat) (*proto.Empty, error) {
	m := Message{
		messageId: req.MessageId,
		data:      req.Data,
	}
	err := s.Impl.OnMessageIn(m)
	return &proto.Empty{}, err
}
