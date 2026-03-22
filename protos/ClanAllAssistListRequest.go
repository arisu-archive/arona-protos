package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ClanAllAssistListRequest struct {
	RequestPacket
	EchelonType          flatdata.EchelonType `json:",omitempty,omitzero"`
	PendingAssistUseInfo []ClanAssistUseInfo
	IsPractice           bool `json:",omitempty,omitzero"`
}
