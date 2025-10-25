package protos

type EventContentTreasureFlipRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
	Round          int32
	Cells          []EventContentTreasureCell
}
