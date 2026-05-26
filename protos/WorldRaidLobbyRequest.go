package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type WorldRaidLobbyRequest struct {
	RequestPacket
	ContentType flatdata.ContentType
	SeasonId    int64 `json:",omitempty,omitzero"`
}
