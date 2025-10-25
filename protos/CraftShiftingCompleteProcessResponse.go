package protos

type CraftShiftingCompleteProcessResponse struct {
	ResponsePacket
	Protocol       Protocol
	CraftInfoDB    ShiftingCraftInfoDB
	ParcelResultDB ParcelResultDB
}
