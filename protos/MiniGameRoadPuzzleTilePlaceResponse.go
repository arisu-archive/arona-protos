package protos

type MiniGameRoadPuzzleTilePlaceResponse struct {
	ResponsePacket
	Protocol        Protocol
	RailTileToPlace RoadPuzzleRailTileData
	SaveDB          RoadPuzzleBoardSaveDB
	ParcelResultDB  ParcelResultDB
}
