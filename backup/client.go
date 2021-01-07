package backup

import (
	"context"
	"test_plugin/proto"
)

// backup client wrap backup service
type BackupClient struct {
	c proto.BackupClient
}

var _ BackupService = (*BackupClient)(nil)

func (c *BackupClient) OnMessageIn(m Message) error {
	_, err := c.c.OnMessageIn(context.TODO(), &proto.Chat{
		MessageId: m.messageId,
		Data:      m.data,
	})
	return err
}
