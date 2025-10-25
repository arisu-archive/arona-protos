package protos

type ShopBuyEligmaRequest struct {
	RequestPacket
	Protocol          Protocol
	GoodsUniqueId     int64
	ShopUniqueId      int64
	CharacterUniqueId int64
	PurchaseCount     int64
}
