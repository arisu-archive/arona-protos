package protos

import (
	"time"
)

type AccountBirthDayRequest struct {
	RequestPacket
	Protocol Protocol
	BirthDay time.Time
}
