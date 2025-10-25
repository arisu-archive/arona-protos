package protos

import (
	"time"
)

type RequestPacket struct {
	BasePacket
	Resendable                    bool
	Hash                          int64
	IsTest                        bool
	ModifiedServerTime__DebugOnly *time.Time
}
