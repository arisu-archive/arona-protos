package protos

type CraftShiftingCompleteProcessRequest struct {
	RequestPacket
	Protocol Protocol
	SlotId   int64
}
