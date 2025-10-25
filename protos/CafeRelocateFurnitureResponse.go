package protos

type CafeRelocateFurnitureResponse struct {
	ResponsePacket
	Protocol             Protocol
	CafeDB               CafeDB
	RelocatedFurnitureDB FurnitureDB
}
