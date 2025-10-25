package protos

type EventContentBoxGachaShopPurchaseResponse struct {
	ResponsePacket
	Protocol               Protocol
	ParcelResultDB         ParcelResultDB
	BoxGachaDB             EventContentBoxGachaDB
	BoxGachaGroupIdByCount map[int64]int64
	BoxGachaElements       []EventContentBoxGachaElement
}
