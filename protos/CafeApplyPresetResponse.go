package protos

type CafeApplyPresetResponse struct {
	ResponsePacket
	Protocol     Protocol
	CafeDBs      []CafeDB
	FurnitureDBs []FurnitureDB
}
