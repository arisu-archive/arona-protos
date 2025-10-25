package protos

type EventContentCardShopPurchaseResponse struct {
	ResponsePacket
	Protocol                   Protocol
	ParcelResultDB             ParcelResultDB
	CardShopElementDB          CardShopElementDB
	CardShopPurchaseHistoryDBs []CardShopPurchaseHistoryDB
}
