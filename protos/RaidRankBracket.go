package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type RaidRankBracket struct {
	Difficulty     flatdata.Difficulty
	BossGroupIndex *int32
	RankCount      int64
}
