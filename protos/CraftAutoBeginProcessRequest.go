package protos

type CraftAutoBeginProcessRequest struct {
	RequestPacket
	PresetSlotDB CraftPresetSlotDB
	Count        int64 `json:",omitempty,omitzero"`
}
