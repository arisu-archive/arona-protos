package protos

type CafeListPresetResponse struct {
	ResponsePacket
	Protocol      Protocol
	CafePresetDBs []CafePresetDB
}
