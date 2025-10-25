package protos

type CafeGetInfoResponse struct {
	ResponsePacket
	Protocol     Protocol
	CafeDB       CafeDB
	CafeDBs      []CafeDB
	FurnitureDBs []FurnitureDB
}
