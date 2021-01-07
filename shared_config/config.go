package SharedConfig

import "github.com/hashicorp/go-plugin"

var HandshakeConfig = &plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BACKUP",
	MagicCookieValue: "FOOBAR",
}
