package protos

type ShopBuyMerchandiseRequest struct {
	RequestPacket
	Protocol       Protocol
	IsRefreshGoods bool
	ShopUniqueId   int64
	GoodsId        int64
	PurchaseCount  int64
}
