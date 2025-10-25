package protos

type CafeRemoveFurnitureResponse struct {
	ResponsePacket
	Protocol     Protocol
	CafeDB       CafeDB
	FurnitureDBs []FurnitureDB
}
