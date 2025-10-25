package protos

type MultiSweepPresetDB struct {
	PresetId   int64
	PresetName string
	StageIds   []int64
	ParcelIds  []ParcelKeyPair
}
