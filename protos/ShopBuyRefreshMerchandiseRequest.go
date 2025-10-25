package protos

type ShopBuyRefreshMerchandiseRequest struct {
	RequestPacket
	Protocol      Protocol
	ShopUniqueIds []int64
}
