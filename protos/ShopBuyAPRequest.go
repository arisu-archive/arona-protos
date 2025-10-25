package protos

type ShopBuyAPRequest struct {
	RequestPacket
	Protocol      Protocol
	ShopUniqueId  int64
	PurchaseCount int64
}
