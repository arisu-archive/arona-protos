package protos

type EventContentBoxGachaShopPurchaseRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
	PurchaseCount  int64
	PurchaseAll    bool
}
