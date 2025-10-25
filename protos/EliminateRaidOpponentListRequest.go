package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type EliminateRaidOpponentListRequest struct {
	RequestPacket
	Protocol       Protocol
	Rank           *int64
	Score          *int64
	BossGroupIndex *int32
	IsUpper        bool
	IsFirstRequest bool
	SearchType     flatdata.RankingSearchType
}
