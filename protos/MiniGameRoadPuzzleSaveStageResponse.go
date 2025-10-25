package protos

type MiniGameRoadPuzzleSaveStageResponse struct {
	ResponsePacket
	Protocol Protocol
	SaveDB   RoadPuzzleBoardSaveDB
}
