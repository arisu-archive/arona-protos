package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type RaidLoginResponse struct {
	ResponsePacket
	Protocol                Protocol
	SeasonType              flatdata.RaidSeasonType
	CanReceiveRankingReward bool
	LastSettledRanking      int64
	LastSettledTier         *int32
}
