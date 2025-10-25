package protos

type ShopBuyGacha3Request struct {
	ShopBuyGacha2Request
	Protocol      Protocol
	FreeRecruitId int64
	Cost          ParcelCost
}
