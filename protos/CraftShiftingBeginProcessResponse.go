package protos

type CraftShiftingBeginProcessResponse struct {
	ResponsePacket
	Protocol       Protocol
	CraftInfoDB    ShiftingCraftInfoDB
	ParcelResultDB ParcelResultDB
}
