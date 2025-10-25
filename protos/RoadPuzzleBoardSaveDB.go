package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type RoadPuzzleBoardSaveDB struct {
	UniqueId             int64
	Round                int32
	RecentRandomRailTile RoadPuzzleRailTileData
	RemainingTiles       map[flatdata.RoadPuzzleRailTileType]int32
	PlacedRailTiles      []RoadPuzzleRailTileData
	RewardTiles          []RoadPuzzleTileRewardData
	IsTrainReadyToDepart bool
}
