package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type RaidLobbyResponse struct {
	ResponsePacket
	SeasonType        flatdata.RaidSeasonType `json:",omitempty,omitzero"`
	RaidGiveUpDB      RaidGiveUpDB
	RaidLobbyInfoDB   SingleRaidLobbyInfoDB
	AccountCurrencyDB AccountCurrencyDB
	ParcelResultDB    ParcelResultDB
}
