package protos

import (
	"time"

	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type CampaignChapterClearRewardHistoryDB struct {
	ChapterUniqueId int64
	RewardType      flatdata.StageDifficulty
	ReceiveDate     time.Time
}
