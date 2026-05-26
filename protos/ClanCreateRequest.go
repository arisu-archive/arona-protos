package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ClanCreateRequest struct {
	RequestPacket
	ClanNickName   string
	ClanJoinOption flatdata.ClanJoinOption `json:",omitempty,omitzero"`
}
