package protos

type EventContentCardShopPurchaseRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
	SlotNumber     int32
}
