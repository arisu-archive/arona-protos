package protos

type EventContentTreasureNextRoundResponse struct {
	ResponsePacket
	Protocol       Protocol
	BoardHistoryDB EventContentTreasureHistoryDB
	HiddenImage    EventContentTreasureCell
}
