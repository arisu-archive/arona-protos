package protos

type ContentSweepSetMultiSweepPresetNameRequest struct {
	RequestPacket
	Protocol   Protocol
	PresetId   int64
	PresetName string
}
