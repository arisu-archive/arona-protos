package protos

import (
	"time"

	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type RaidLobbyInfoDB struct {
	SeasonId                      int64
	Tier                          int32
	Ranking                       int64
	BestRankingPoint              int64
	TotalRankingPoint             int64
	ReceivedRankingRewardId       int64
	CanReceiveRankingReward       bool
	PlayingRaidDB                 RaidDB
	ReceiveRewardIds              []int64
	ReceiveLimitedRewardIds       []int64
	ParticipateCharacterServerIds []int64
	PlayableHighestDifficulty     map[string]flatdata.Difficulty
	SweepPointByRaidUniqueId      map[int64]int64
	SeasonStartDate               time.Time
	SeasonEndDate                 time.Time
	SettlementEndDate             time.Time
	NextSeasonId                  int64
	NextSeasonStartDate           time.Time
	NextSeasonEndDate             time.Time
	NextSettlementEndDate         time.Time
	ClanAssistUseInfo             ClanAssistUseInfo
	RemainFailCompensation        map[int32]bool
}
