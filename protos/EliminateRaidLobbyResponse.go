package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type EliminateRaidLobbyResponse struct {
	ResponsePacket
	Protocol          Protocol
	SeasonType        flatdata.RaidSeasonType
	RaidGiveUpDB      RaidGiveUpDB
	RaidLobbyInfoDB   EliminateRaidLobbyInfoDB
	AccountCurrencyDB AccountCurrencyDB
	ParcelResultDB    ParcelResultDB
}
