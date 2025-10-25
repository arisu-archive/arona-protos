package protos

type CafeRenamePresetRequest struct {
	RequestPacket
	Protocol   Protocol
	SlotId     int32
	PresetName string
}
