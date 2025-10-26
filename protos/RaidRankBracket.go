package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type RaidRankBracket struct {
	Difficulty flatdata.Difficulty `json:",omitempty,omitzero"`
	BossGroupIndex *int32 `json:",omitempty,omitzero"`
	RankCount int64 `json:",omitempty,omitzero"`
}
