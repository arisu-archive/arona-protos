package protos

type ContentSweepMultiSweepPresetListResponse struct {
	ResponsePacket
	Protocol            Protocol
	MultiSweepPresetDBs []MultiSweepPresetDB
}
