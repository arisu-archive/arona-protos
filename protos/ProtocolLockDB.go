package protos

import (
	"time"
)

type ProtocolLockDB struct {
	ProtocolId int32 `json:",omitempty,omitzero"`
	BeginDate time.Time `json:",omitempty,omitzero"`
	EndDate time.Time `json:",omitempty,omitzero"`
	CreateDate time.Time `json:",omitempty,omitzero"`
}
