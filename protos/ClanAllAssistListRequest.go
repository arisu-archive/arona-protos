package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ClanAllAssistListRequest struct {
	RequestPacket
	Protocol             Protocol
	EchelonType          flatdata.EchelonType
	PendingAssistUseInfo []ClanAssistUseInfo
	IsPractice           bool
}
