package protos

type CraftInfoListResponse struct {
	ResponsePacket
	Protocol           Protocol
	CraftInfos         []CraftInfoDB
	ShiftingCraftInfos []ShiftingCraftInfoDB
}
