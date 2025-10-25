package protos

import (
	"time"
)

type ArenaHistoryRequest struct {
	RequestPacket
	Protocol        Protocol
	SearchStartDate time.Time
	Count           int32
}
