package protos

type EventContentCardShopPurchaseAllResponse struct {
	ResponsePacket
	Protocol                   Protocol
	ParcelResultDB             ParcelResultDB
	CardShopElementDBs         []CardShopElementDB
	CardShopPurchaseHistoryDBs []CardShopPurchaseHistoryDB
	RewardHistory              map[int64][]ParcelInfo
}
