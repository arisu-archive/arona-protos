package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ClanCreateRequest struct {
	RequestPacket
	Protocol       Protocol
	ClanNickName   string
	ClanJoinOption flatdata.ClanJoinOption
}
