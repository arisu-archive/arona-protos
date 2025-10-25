package protos

type CraftShiftingCompleteProcessAllResponse struct {
	ResponsePacket
	Protocol       Protocol
	CraftInfoDBs   []ShiftingCraftInfoDB
	ParcelResultDB ParcelResultDB
}
