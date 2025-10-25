package protos

type EventContentMainGroundStageResultRequest struct {
	RequestPacket
	Protocol           Protocol
	EventContentId     int64
	PassCheckCharacter bool
	Summary            BattleSummary
}
