package protos

type EventContentTreasureFlipResponse struct {
	ResponsePacket
	Protocol       Protocol
	BoardHistoryDB EventContentTreasureHistoryDB
	ParcelResultDB ParcelResultDB
}
