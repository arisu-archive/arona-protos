package protos

import (
	"time"

	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type RaidSeasonRankingHistoryDB struct {
	ContentType      flatdata.ContentType
	AccountId        int64
	SeasonId         int64
	Ranking          int64
	BestRankingPoint int64
	Tier             int32
	ReceivedDate     time.Time
}
