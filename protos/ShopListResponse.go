package protos

type ShopListResponse struct {
	ResponsePacket
	Protocol             Protocol
	ShopInfos            []ShopInfoDB
	ShopEligmaHistoryDBs []ShopEligmaHistoryDB
}
