package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ClanSettingRequest struct {
	RequestPacket
	ChangedClanName string
	ChangedNotice   string
	ClanJoinOption  flatdata.ClanJoinOption `json:",omitempty,omitzero"`
}
