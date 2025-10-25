package protos

type EventContentShopListResponse struct {
	ResponsePacket
	Protocol             Protocol
	ShopInfos            []ShopInfoDB
	ShopEligmaHistoryDBs []ShopEligmaHistoryDB
}
