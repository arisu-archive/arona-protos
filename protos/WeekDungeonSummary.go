package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type WeekDungeonSummary struct {
	DungeonType flatdata.WeekDungeonType
	FindGifts   []FindGiftSummary
}
