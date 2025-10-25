package protos

type MiniGameRoadPuzzleTilePlaceRequest struct {
	RequestPacket
	Protocol        Protocol
	EventContentId  int64
	UniqueId        int64
	Round           int64
	RailTileToPlace RoadPuzzleRailTileData
}
