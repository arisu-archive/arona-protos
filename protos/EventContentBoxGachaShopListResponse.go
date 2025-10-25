package protos

type EventContentBoxGachaShopListResponse struct {
	ResponsePacket
	Protocol               Protocol
	BoxGachaDB             EventContentBoxGachaDB
	BoxGachaGroupIdByCount map[int64]int64
}
