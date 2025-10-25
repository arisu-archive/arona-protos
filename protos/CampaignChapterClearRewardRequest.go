package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type CampaignChapterClearRewardRequest struct {
	RequestPacket
	Protocol                Protocol
	CampaignChapterUniqueId int64
	StageDifficulty         flatdata.StageDifficulty
}
