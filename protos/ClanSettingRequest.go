package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ClanSettingRequest struct {
	RequestPacket
	Protocol        Protocol
	ChangedClanName string
	ChangedNotice   string
	ClanJoinOption  flatdata.ClanJoinOption
}
