package protos

type CraftBeginProcessRequest struct {
	RequestPacket
	Protocol Protocol
	SlotId   int64
}
