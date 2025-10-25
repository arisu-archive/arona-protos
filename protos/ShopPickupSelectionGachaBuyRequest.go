package protos

type ShopPickupSelectionGachaBuyRequest struct {
	ShopBuyGacha2Request
	Protocol      Protocol
	FreeRecruitId int64
	Cost          ParcelCost
}
