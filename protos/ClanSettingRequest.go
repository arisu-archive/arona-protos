package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ClanSettingRequest struct {
	RequestPacket
	ChangedClanName string                  `json:",omitempty,omitzero"`
	ChangedNotice   string                  `json:",omitempty,omitzero"`
	ClanJoinOption  flatdata.ClanJoinOption `json:",omitempty,omitzero"`
}
