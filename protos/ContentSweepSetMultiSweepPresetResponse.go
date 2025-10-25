package protos

type ContentSweepSetMultiSweepPresetResponse struct {
	ResponsePacket
	Protocol            Protocol
	MultiSweepPresetDBs []MultiSweepPresetDB
}
