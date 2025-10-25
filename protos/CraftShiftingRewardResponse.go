package protos

type CraftShiftingRewardResponse struct {
	ResponsePacket
	Protocol         Protocol
	ParcelResultDB   ParcelResultDB
	TargetCraftInfos []ShiftingCraftInfoDB
}
