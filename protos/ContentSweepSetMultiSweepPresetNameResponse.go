package protos

type ContentSweepSetMultiSweepPresetNameResponse struct {
	ResponsePacket
	Protocol            Protocol
	MultiSweepPresetDBs []MultiSweepPresetDB
}
