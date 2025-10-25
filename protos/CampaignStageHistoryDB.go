package protos

import (
	"time"
)

type CampaignStageHistoryDB struct {
	StoryUniqueId                   int64
	ChapterUniqueId                 int64
	StageUniqueId                   int64
	TacticClearCountWithRankSRecord int64
	ClearTurnRecord                 int64
	Star1Flag                       bool
	Star2Flag                       bool
	Star3Flag                       bool
	LastPlay                        time.Time
	TodayPlayCount                  int64
	TodayPurchasePlayCountHardStage int64
	FirstClearRewardReceive         *time.Time
	StarRewardReceive               *time.Time
	TodayPlayCountForUI             int64
}
