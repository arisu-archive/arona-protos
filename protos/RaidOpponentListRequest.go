package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type RaidOpponentListRequest struct {
	RequestPacket
	Rank           *int64                     `json:",omitempty,omitzero"`
	Score          *int64                     `json:",omitempty,omitzero"`
	IsUpper        bool                       `json:",omitempty,omitzero"`
	IsFirstRequest bool                       `json:",omitempty,omitzero"`
	SearchType     flatdata.RankingSearchType `json:",omitempty,omitzero"`
}
