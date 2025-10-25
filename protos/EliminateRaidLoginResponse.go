package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type EliminateRaidLoginResponse struct {
	ResponsePacket
	Protocol                 Protocol
	SeasonType               flatdata.RaidSeasonType
	CanReceiveRankingReward  bool
	ReceiveLimitedRewardIds  []int64
	SweepPointByRaidUniqueId map[int64]int64
	LastSettledRanking       int64
	LastSettledTier          *int32
}
