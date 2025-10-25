package protos

type MiniGameRoadPuzzleGetInfoResponse struct {
	ResponsePacket
	Protocol Protocol
	SaveDB   RoadPuzzleBoardSaveDB
}
