package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type RaidLobbyResponse struct {
	ResponsePacket
	Protocol          Protocol
	SeasonType        flatdata.RaidSeasonType
	RaidGiveUpDB      RaidGiveUpDB
	RaidLobbyInfoDB   SingleRaidLobbyInfoDB
	AccountCurrencyDB AccountCurrencyDB
	ParcelResultDB    ParcelResultDB
}
