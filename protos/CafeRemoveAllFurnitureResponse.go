package protos

type CafeRemoveAllFurnitureResponse struct {
	ResponsePacket
	Protocol     Protocol
	CafeDB       CafeDB
	FurnitureDBs []FurnitureDB
}
