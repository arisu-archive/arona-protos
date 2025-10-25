package protos

type CraftRewardRequest struct {
	RequestPacket
	Protocol Protocol
	SlotId   int64
}
