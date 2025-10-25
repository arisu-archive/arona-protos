package protos

type EventContentTreasureLobbyResponse struct {
	ResponsePacket
	Protocol       Protocol
	BoardHistoryDB EventContentTreasureHistoryDB
	HiddenImage    EventContentTreasureCell
	VariationId    int64
}
