package protos

import (
	"time"
)

type ClanChatLogRequest struct {
	RequestPacket
	Protocol Protocol
	Channel  string
	FromDate time.Time
}
