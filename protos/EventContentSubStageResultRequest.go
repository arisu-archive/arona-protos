package protos

type EventContentSubStageResultRequest struct {
	RequestPacket
	Protocol           Protocol
	EventContentId     int64
	PassCheckCharacter bool
	Summary            BattleSummary
}
