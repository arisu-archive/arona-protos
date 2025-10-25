package protos

type CraftCompleteProcessRequest struct {
	RequestPacket
	Protocol Protocol
	SlotId   int64
}
