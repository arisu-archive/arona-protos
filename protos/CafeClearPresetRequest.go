package protos

type CafeClearPresetRequest struct {
	RequestPacket
	Protocol Protocol
	SlotId   int32
}
