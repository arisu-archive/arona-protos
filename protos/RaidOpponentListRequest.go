package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type RaidOpponentListRequest struct {
	RequestPacket
	Protocol       Protocol
	Rank           *int64
	Score          *int64
	IsUpper        bool
	IsFirstRequest bool
	SearchType     flatdata.RankingSearchType
}
