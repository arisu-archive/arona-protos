package protos

type ItemConsumeResponse struct {
	ResponsePacket
	Protocol          Protocol
	UsedItemDB        ItemDB
	NewParcelResultDB ParcelResultDB
}
