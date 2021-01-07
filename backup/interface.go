package backup

type BackupService interface {
	OnMessageIn(message Message) error
}

// message for backup, this should be on other package btw
type Message struct {
	messageId string
	data      string
}
