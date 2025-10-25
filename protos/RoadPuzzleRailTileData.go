package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type RoadPuzzleRailTileData struct {
	RoadPuzzleTileData
	Type          flatdata.RoadPuzzleRailTileType
	EntranceIndex int32
	ExitIndex     int32
	ResourcePath  string
}
