package protos

type CraftAutoBeginProcessRequest struct {
	RequestPacket
	Protocol     Protocol
	PresetSlotDB CraftPresetSlotDB
	Count        int64
}
