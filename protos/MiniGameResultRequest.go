package protos

type MiniGameResultRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
	UniqueId       int64
	Summary        MinigameRhythmSummary
}
