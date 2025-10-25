package protos

import (
	"time"

	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ConquestTileDB struct {
	EventContentId   int64
	Difficulty       flatdata.StageDifficulty
	TileUniqueId     int64
	TileState        flatdata.TileState
	Level            int64
	StarFlags        []bool
	CreateTime       time.Time
	IsThreeStarClear bool
	IsAnyStarClear   bool
}
