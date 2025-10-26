package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ClanSearchRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SearchString string `json:",omitempty,omitzero"`
	ClanJoinOption flatdata.ClanJoinOption `json:",omitempty,omitzero"`
	ClanUniqueCode string `json:",omitempty,omitzero"`
}
