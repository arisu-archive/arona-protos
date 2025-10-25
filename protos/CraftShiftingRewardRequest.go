package protos

type CraftShiftingRewardRequest struct {
	RequestPacket
	Protocol Protocol
	SlotId   int64
}
