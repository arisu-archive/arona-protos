package protos

import (
	"time"
)

type ProtocolLockDB struct {
	ProtocolId int32
	BeginDate  time.Time
	EndDate    time.Time
	CreateDate time.Time
}
