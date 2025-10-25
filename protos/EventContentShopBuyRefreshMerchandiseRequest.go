package protos

type EventContentShopBuyRefreshMerchandiseRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
	ShopUniqueIds  []int64
}
