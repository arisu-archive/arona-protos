package protos

type EventContentTreasureNextRoundRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
	Round          int32
}
