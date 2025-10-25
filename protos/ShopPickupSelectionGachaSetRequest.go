package protos

type ShopPickupSelectionGachaSetRequest struct {
	RequestPacket
	Protocol                 Protocol
	ShopRecruitID            int64
	PickupCharacterSelection map[int64]int64
}
