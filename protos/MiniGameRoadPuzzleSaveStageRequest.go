package protos

type MiniGameRoadPuzzleSaveStageRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
	UniqueId       int64
	Round          int64
	PlaceRailTiles []RoadPuzzleRailTileData
}
