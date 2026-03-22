package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ClanSearchRequest struct {
	RequestPacket
	SearchString   string
	ClanJoinOption flatdata.ClanJoinOption `json:",omitempty,omitzero"`
	ClanUniqueCode string
}
