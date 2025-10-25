package protos

type CraftShiftingRewardAllResponse struct {
	ResponsePacket
	Protocol       Protocol
	ParcelResultDB ParcelResultDB
	CraftInfoDBs   []ShiftingCraftInfoDB
}
