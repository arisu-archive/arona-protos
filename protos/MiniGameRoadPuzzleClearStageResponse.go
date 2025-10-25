package protos

type MiniGameRoadPuzzleClearStageResponse struct {
	ResponsePacket
	Protocol       Protocol
	IsSkip         bool
	ParcelResultDB ParcelResultDB
}
