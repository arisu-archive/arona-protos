package protos

type ShopPickupSelectionGachaGetResponse struct {
	ResponsePacket
	Protocol                 Protocol
	PickupCharacterSelection map[int64]int64
}
