package protos

import (
	"time"
)

type AccountRequestBirthdayMailRequest struct {
	RequestPacket
	Protocol Protocol
	Birthday time.Time
}
