package protos

type MiniGameRoadPuzzleClearStageRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
	UniqueId       int64
	Round          int64
	IsSkip         bool
}
