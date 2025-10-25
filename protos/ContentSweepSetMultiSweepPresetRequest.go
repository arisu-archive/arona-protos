package protos

type ContentSweepSetMultiSweepPresetRequest struct {
	RequestPacket
	Protocol   Protocol
	PresetId   int64
	PresetName string
	StageIds   []int64
	ParcelIds  []ParcelKeyPair
}
