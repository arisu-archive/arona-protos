package protos

type CafeApplyTemplateResponse struct {
	ResponsePacket
	Protocol     Protocol
	CafeDBs      []CafeDB
	FurnitureDBs []FurnitureDB
}
