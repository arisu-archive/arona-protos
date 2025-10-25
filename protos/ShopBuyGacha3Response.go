package protos

type ShopBuyGacha3Response struct {
	ShopBuyGacha2Response
	Protocol                 Protocol
	FreeRecruitHistoryDB     ShopFreeRecruitHistoryDB
	PickupFirstGetHistoryDBs []PickupFirstGetHistoryDB
}
