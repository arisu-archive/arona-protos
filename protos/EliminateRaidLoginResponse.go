package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
	"github.com/arisu-archive/mapx"
)

type EliminateRaidLoginResponse struct {
	ResponsePacket
	SeasonType               flatdata.RaidSeasonType `json:",omitempty,omitzero"`
	CanReceiveRankingReward  bool                    `json:",omitempty,omitzero"`
	ReceiveLimitedRewardIds  []int64
	SweepPointByRaidUniqueId *mapx.OrderedMap[int64, int64]
	LastSettledRanking       int64  `json:",omitempty,omitzero"`
	LastSettledTier          *int32 `json:",omitempty,omitzero"`
}
