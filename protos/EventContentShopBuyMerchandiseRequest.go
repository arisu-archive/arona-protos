package protos

type EventContentShopBuyMerchandiseRequest struct {
	RequestPacket
	Protocol             Protocol
	EventContentId       int64
	IsRefreshMerchandise bool
	ShopUniqueId         int64
	GoodsUniqueId        int64
	PurchaseCount        int64
}
